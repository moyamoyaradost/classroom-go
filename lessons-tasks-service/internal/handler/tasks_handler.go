package handler

import (
    "context"
    "database/sql"
    "github.com/google/uuid"
    pb "github.com/moyamoyaradost/classroom-go/lessons-tasks-service/api/v1"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type TasksServer struct {
    pb.UnimplementedTasksServiceServer
    db *sql.DB
}

func NewTasksServer(db *sql.DB) *TasksServer {
    return &TasksServer{db: db}
}

func (s *TasksServer) CreateTask(ctx context.Context, req *pb.TaskCreateRequest) (*pb.TaskResponse, error) {
    id := uuid.New().String()
    
    _, err := s.db.ExecContext(ctx,
        `INSERT INTO tasks (id, course_id, description) 
        VALUES ($1, $2, $3)`,
        id, req.CourseId, req.Description,
    )
    
    if err != nil {
        return nil, status.Error(codes.Internal, "failed to create task")
    }

    return &pb.TaskResponse{
        Id:          id,
        CourseId:    req.CourseId,
        Description: req.Description,
    }, nil
}

func (s *TasksServer) GetTasksByCourse(ctx context.Context, req *pb.CourseIdRequest) (*pb.TasksListResponse, error) {
    // Заглушка для нереализованного метода
    return nil, status.Error(codes.Unimplemented, "method GetTasksByCourse not implemented")
}

func (s *TasksServer) CompleteTask(ctx context.Context, req *pb.UserTaskRequest) (*pb.CompleteTaskResponse, error) {
    // Заглушка для нереализованного метода
    return nil, status.Error(codes.Unimplemented, "method CompleteTask not implemented")
}
