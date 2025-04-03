package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	v1 "github.com/moyamoyaradost/classroom-go/lessons-tasks-service/api/v1"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/internal/handler"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/internal/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting lessons-tasks service...")

	// Получение переменных окружения
	dbHost := getEnv("DB_HOST", "postgres")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "secret")
	dbName := getEnv("DB_NAME", "classroom")
	redisAddr := getEnv("REDIS_ADDR", "redis:6379")
	grpcPort := getEnv("GRPC_PORT", "50051")

	// Инициализация подключения к PostgreSQL
	dbConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Проверка соединения с БД
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Connected to PostgreSQL")

	// Настройка Redis для кэширования
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	defer rdb.Close()

	// Проверка соединения с Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")

	// Инициализация репозиториев
	lessonRepo := repository.NewLessonRepo(db)
	taskRepo := repository.NewTaskRepo(db)
	userTaskRepo := repository.NewUserTaskRepo(db)
	answerRepo := repository.NewTaskAnswerRepo(db) // Добавленный репозиторий

	// Инициализация обработчиков
	lessonServer := handler.NewLessonsServer(lessonRepo)
	taskServer := handler.NewTasksServer(taskRepo, userTaskRepo, answerRepo) // Исправленный вызов

	// Запуск gRPC-сервера
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Регистрация сервисов
	v1.RegisterLessonsServiceServer(grpcServer, lessonServer)
	v1.RegisterTasksServiceServer(grpcServer, taskServer)

	// Включение рефлексии для инструментов отладки
	reflection.Register(grpcServer)

	// Запуск сервера в горутине
	go func() {
		log.Printf("Starting gRPC server on port %s", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Обработка сигналов для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
