package handler

import (
	"context"
	"time"

	"github.com/google/uuid"
	v1 "github.com/moyamoyaradost/classroom-go/lessons-tasks-service/api/v1"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/internal/repository"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TasksServer struct {
	v1.UnimplementedTasksServiceServer
	taskRepo     *repository.TaskRepo
	userTaskRepo *repository.UserTaskRepo
	answerRepo   *repository.TaskAnswerRepo
}

func NewTasksServer(taskRepo *repository.TaskRepo, userTaskRepo *repository.UserTaskRepo, answerRepo *repository.TaskAnswerRepo) *TasksServer {
	return &TasksServer{
		taskRepo:     taskRepo,
		userTaskRepo: userTaskRepo,
		answerRepo:   answerRepo,
	}
}

func (s *TasksServer) SubmitTaskAnswer(ctx context.Context, req *v1.SubmitAnswerRequest) (*v1.SubmitAnswerResponse, error) {
	if _, err := s.taskRepo.GetTaskByID(ctx, req.TaskId); err != nil {
		return nil, status.Error(codes.NotFound, "задание не найдено")
	}

	answer := models.TaskAnswer{
		ID:             uuid.New().String(),
		TaskID:         req.TaskId,
		UserID:         req.UserId,
		AnswerText:     req.AnswerText,
		AttachmentURLs: req.AttachmentUrls,
		Status:         v1.AnswerStatus_PENDING.String(), // Преобразуем enum в строку
		CreatedAt:      time.Now(),
	}

	if err := s.answerRepo.SaveTaskAnswer(ctx, answer); err != nil {
		return nil, status.Error(codes.Internal, "ошибка сохранения ответа")
	}

	// Преобразование времени в protobuf-формат
	createdAtProto := timestamppb.New(answer.CreatedAt)

	return &v1.SubmitAnswerResponse{
		Id:             answer.ID,
		TaskId:         answer.TaskID,
		UserId:         answer.UserID,
		AnswerText:     answer.AnswerText,
		AttachmentUrls: answer.AttachmentURLs,
		Status:         v1.AnswerStatus(v1.AnswerStatus_value[answer.Status]), // Конвертация строки в enum
		CreatedAt:      createdAtProto,
	}, nil
}

func (s *TasksServer) GetTaskAnswers(ctx context.Context, req *v1.TaskAnswersRequest) (*v1.TaskAnswersResponse, error) {
	var answers []models.TaskAnswer
	var err error

	if req.UserId != "" {
		answers, err = s.answerRepo.GetTaskAnswersByUser(ctx, req.TaskId, req.UserId)
	} else {
		answers, err = s.answerRepo.GetAllTaskAnswers(ctx, req.TaskId)
	}

	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка при получении ответов на задание")
	}

	response := &v1.TaskAnswersResponse{
		Answers: make([]*v1.SubmitAnswerResponse, 0, len(answers)),
	}

	for _, answer := range answers {
		// Преобразование времени
		createdAtProto := timestamppb.New(answer.CreatedAt)
		
		// Конвертация статуса из строки в enum
		statusEnum := v1.AnswerStatus(v1.AnswerStatus_value[answer.Status])

		response.Answers = append(response.Answers, &v1.SubmitAnswerResponse{
			Id:             answer.ID,
			TaskId:         answer.TaskID,
			UserId:         answer.UserID,
			AnswerText:     answer.AnswerText,
			AttachmentUrls: answer.AttachmentURLs,
			Status:         statusEnum,
			CreatedAt:      createdAtProto,
		})
	}

	return response, nil
}

// Остальные методы остаются без изменений
