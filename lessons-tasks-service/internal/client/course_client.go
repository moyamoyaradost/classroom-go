package client

import "context"

type CourseClient struct{}

func NewCourseClient() *CourseClient {
    return &CourseClient{}
}

// Заглушка для тестирования
func (c *CourseClient) IsTeacher(ctx context.Context, userID, courseID string) (bool, error) {
    return true, nil // Замените на реальную реализацию
}