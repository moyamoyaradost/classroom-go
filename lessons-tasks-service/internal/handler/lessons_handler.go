package handler

import (
    "context"
    pb "github.com/moyamoyaradost/classroom-go/lessons-tasks-service/api/v1"
)

type LessonsServer struct {
    pb.UnimplementedLessonsServiceServer
}

func (s *LessonsServer) CreateLesson(ctx context.Context, req *pb.LessonCreateRequest) (*pb.LessonResponse, error) {
    return &pb.LessonResponse{
        Id:          "generated-id",
        CourseId:    req.CourseId,
        Description: req.Description,
    }, nil
}
