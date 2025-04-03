package handler

import (
    "context"
    "database/sql"
    "github.com/google/uuid"
    pb "github.com/moyamoyaradost/classroom-go/lessons-tasks-service/api/v1"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type LessonsServer struct {
    pb.UnimplementedLessonsServiceServer
    db *sql.DB
}

func NewLessonsServer(db *sql.DB) *LessonsServer {
    return &LessonsServer{db: db}
}

func (s *LessonsServer) CreateLesson(ctx context.Context, req *pb.LessonCreateRequest) (*pb.LessonResponse, error) {
    id := uuid.New().String()
    
    _, err := s.db.ExecContext(ctx,
        `INSERT INTO lessons (id, course_id, description) 
        VALUES ($1, $2, $3)`,
        id, req.CourseId, req.Description,
    )
    
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
    // Заглушка для нереализованного метода
    return nil, status.Error(codes.Unimplemented, "method GetLessonsByCourse not implemented")
}

func (s *LessonsServer) DeleteLesson(ctx context.Context, req *pb.LessonIdRequest) (*pb.DeleteResponse, error) {
    // Заглушка для нереализованного метода
    return nil, status.Error(codes.Unimplemented, "method DeleteLesson not implemented")
}
