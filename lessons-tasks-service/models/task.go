// lessons-tasks-service/models/task.go
package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
