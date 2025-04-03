package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/internal/repository"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/models"
	"github.com/stretchr/testify/assert"
)

func TestTaskRepo_CreateTask(t *testing.T) {
	// Настройка мок-базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Создание репозитория с мок-базой
	repo := repository.NewTaskRepo(db)

	// Подготовка тестовых данных
	task := models.Task{
		ID:          "task-1",
		CourseID:    "course-1",
		Description: "Test task description",
	}

	// Настройка ожидаемого запроса
	mock.ExpectExec("INSERT INTO tasks").
		WithArgs(task.ID, task.CourseID, task.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Выполнение тестируемого метода
	err = repo.CreateTask(context.Background(), task)

	// Проверки
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTaskRepo_GetTaskByID(t *testing.T) {
	// Настройка мок-базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Создание репозитория с мок-базой
	repo := repository.NewTaskRepo(db)

	// Подготовка тестовых данных
	taskID := "task-1"
	courseID := "course-1"
	description := "Test task description"
	createdAt := time.Now()

	// Настройка ожидаемого запроса
	rows := sqlmock.NewRows([]string{"id", "course_id", "description", "created_at"}).
		AddRow(taskID, courseID, description, createdAt)

	mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id = \\$1").
		WithArgs(taskID).
		WillReturnRows(rows)

	// Выполнение тестируемого метода
	task, err := repo.GetTaskByID(context.Background(), taskID)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, taskID, task.ID)
	assert.Equal(t, courseID, task.CourseID)
	assert.Equal(t, description, task.Description)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserTaskRepo_CompleteTask(t *testing.T) {
	// Настройка мок-базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Создание репозитория с мок-базой
	repo := repository.NewUserTaskRepo(db)

	// Подготовка тестовых данных
	userID := "user-1"
	taskID := "task-1"

	// Настройка ожидаемого запроса
	mock.ExpectExec("INSERT INTO user_tasks").
		WithArgs(userID, taskID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Выполнение тестируемого метода
	err = repo.CompleteTask(context.Background(), userID, taskID)

	// Проверки
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserTaskRepo_GetCompletedTasks(t *testing.T) {
	// Настройка мок-базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Создание репозитория с мок-базой
	repo := repository.NewUserTaskRepo(db)

	// Подготовка тестовых данных
	userID := "user-1"
	createdAt := time.Now()
	tasks := []models.Task{
		{ID: "task-1", CourseID: "course-1", Description: "Task 1", CreatedAt: createdAt},
		{ID: "task-2", CourseID: "course-1", Description: "Task 2", CreatedAt: createdAt},
	}

	// Настройка ожидаемого запроса
	rows := sqlmock.NewRows([]string{"id", "course_id", "description", "created_at"})
	for _, t := range tasks {
		rows.AddRow(t.ID, t.CourseID, t.Description, t.CreatedAt)
	}

	mock.ExpectQuery("SELECT t.id, t.course_id, t.description, t.created_at FROM tasks t JOIN user_tasks ut ON t.id = ut.task_id WHERE ut.user_id = \\$1").
		WithArgs(userID).
		WillReturnRows(rows)

	// Выполнение тестируемого метода
	result, err := repo.GetCompletedTasks(context.Background(), userID)

	// Проверки
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "task-1", result[0].ID)
	assert.Equal(t, "task-2", result[1].ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
