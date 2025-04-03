package handler

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/moyamoyaradost/classroom-go/lessons-tasks-service/api/v1"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/internal/repository"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LessonsServer struct {
	pb.UnimplementedLessonsServiceServer
	repo *repository.LessonRepo
}

func NewLessonsServer(repo *repository.LessonRepo) *LessonsServer {
	return &LessonsServer{repo: repo}
}

func (s *LessonsServer) CreateLesson(ctx context.Context, req *pb.LessonCreateRequest) (*pb.LessonResponse, error) {
	id := uuid.New().String()

	// Создаем модель урока
	lesson := models.Lesson{
		ID:          id,
		CourseID:    req.CourseId,
		Description: req.Description,
	}

	// Используем репозиторий вместо прямого доступа к БД
	err := s.repo.CreateLesson(ctx, lesson)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create lesson")
	}

	return &pb.LessonResponse{
		Id:          id,
		CourseId:    req.CourseId,
		Description: req.Description,
	}, nil
}

func (s *LessonsServer) GetLessonsByCourse(ctx context.Context, req *pb.CourseIdRequest) (*pb.LessonsListResponse, error) {
	// Используем репозиторий для получения уроков
	lessons, err := s.repo.GetLessonsByCourse(ctx, req.CourseId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get lessons")
	}

	response := &pb.LessonsListResponse{}
	for _, lesson := range lessons {
		response.Lessons = append(response.Lessons, &pb.LessonResponse{
			Id:          lesson.ID,
			CourseId:    lesson.CourseID,
			Description: lesson.Description,
		})
	}

	return response, nil
}

func (s *LessonsServer) DeleteLesson(ctx context.Context, req *pb.LessonIdRequest) (*pb.DeleteResponse, error) {
	// Используем репозиторий для удаления урока
	err := s.repo.DeleteLesson(ctx, req.LessonId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to delete lesson")
	}

	return &pb.DeleteResponse{
		Success: true,
	}, nil
}
