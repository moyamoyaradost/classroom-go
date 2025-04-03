// lessons-tasks-service/models/lesson.go
package models

import "time"

type Lesson struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
