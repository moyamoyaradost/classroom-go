package middleware

import (
    "context"
    "github.com/moyamoyaradost/classroom-go/lessons-tasks-service/internal/client"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
)

var courseClient *client.CourseClient

func init() {
    courseClient = client.NewCourseClient() // Инициализация клиента
}

// IsTeacher проверяет права доступа
func IsTeacher(ctx context.Context, courseID string) error {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return status.Error(codes.Unauthenticated, "метаданные не найдены")
    }

    userIDs := md.Get("user-id")
    if len(userIDs) == 0 {
        return status.Error(codes.Unauthenticated, "идентификатор пользователя не найден")
    }

    isTeacher, err := courseClient.IsTeacher(ctx, userIDs[0], courseID)
    if err != nil {
        return status.Error(codes.Internal, "ошибка проверки прав доступа")
    }

    if !isTeacher {
        return status.Error(codes.PermissionDenied, "требуются права преподавателя")
    }

    return nil
}
