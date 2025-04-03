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

func (s *TasksServer) CreateTask(ctx context.Context, req *v1.TaskCreateRequest) (*v1.TaskResponse, error) {
	task := models.Task{
		ID:          uuid.New().String(),
		CourseID:    req.CourseId,
		Description: req.Description,
	}

	if err := s.taskRepo.CreateTask(ctx, task); err != nil {
		return nil, status.Error(codes.Internal, "ошибка при создании задания")
	}

	return &v1.TaskResponse{
		Id:          task.ID,
		CourseId:    task.CourseID,
		Description: task.Description,
	}, nil
}

func (s *TasksServer) GetTasksByCourse(ctx context.Context, req *v1.CourseIdRequest) (*v1.TasksListResponse, error) {
	tasks, err := s.taskRepo.GetTasksByCourse(ctx, req.CourseId)
	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка при получении заданий")
	}

	response := &v1.TasksListResponse{
		Tasks: make([]*v1.TaskResponse, 0, len(tasks)),
	}

	for _, task := range tasks {
		response.Tasks = append(response.Tasks, &v1.TaskResponse{
			Id:          task.ID,
			CourseId:    task.CourseID,
			Description: task.Description,
		})
	}

	return response, nil
}

func (s *TasksServer) CompleteTask(ctx context.Context, req *v1.UserTaskRequest) (*v1.CompleteTaskResponse, error) {
	err := s.userTaskRepo.CompleteTask(ctx, req.UserId, req.TaskId)
	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка при выполнении задания")
	}

	return &v1.CompleteTaskResponse{
		Success: true,
	}, nil
}

func (s *TasksServer) GetCompletedTasksByUser(ctx context.Context, req *v1.UserIdRequest) (*v1.UserTasksListResponse, error) {
	tasks, err := s.userTaskRepo.GetCompletedTasks(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "ошибка при получении выполненных заданий")
	}

	response := &v1.UserTasksListResponse{
		Tasks: make([]*v1.TaskResponse, 0, len(tasks)),
	}

	for _, task := range tasks {
		response.Tasks = append(response.Tasks, &v1.TaskResponse{
			Id:          task.ID,
			CourseId:    task.CourseID,
			Description: task.Description,
		})
	}

	return response, nil
}

func (s *TasksServer) SubmitTaskAnswer(ctx context.Context, req *v1.SubmitAnswerRequest) (*v1.SubmitAnswerResponse, error) {
	// Убрана неиспользуемая переменная task
	if _, err := s.taskRepo.GetTaskByID(ctx, req.TaskId); err != nil {
		return nil, status.Error(codes.NotFound, "задание не найдено")
	}

	answer := models.TaskAnswer{
		ID:             uuid.New().String(),
		TaskID:         req.TaskId,
		UserID:         req.UserId,
		AnswerText:     req.AnswerText,
		AttachmentURLs: req.AttachmentUrls,
		Status:         "pending",
		CreatedAt:      time.Now(),
	}

	if err := s.answerRepo.SaveTaskAnswer(ctx, answer); err != nil {
		return nil, status.Error(codes.Internal, "ошибка сохранения ответа")
	}

	return &v1.SubmitAnswerResponse{
		Id:             answer.ID,
		TaskId:         answer.TaskID,
		UserId:         answer.UserID,
		AnswerText:     answer.AnswerText,
		AttachmentUrls: answer.AttachmentURLs,
		Status:         answer.Status,
		CreatedAt:      answer.CreatedAt.Format(time.RFC3339),
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
		response.Answers = append(response.Answers, &v1.SubmitAnswerResponse{
			Id:             answer.ID,
			TaskId:         answer.TaskID,
			UserId:         answer.UserID,
			AnswerText:     answer.AnswerText,
			AttachmentUrls: answer.AttachmentURLs,
			Status:         answer.Status,
			CreatedAt:      answer.CreatedAt.Format(time.RFC3339),
		})
	}

	return response, nil
}

