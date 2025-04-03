// lessons-tasks-service/models/user_task_model.go
package models

import "time"

// UserTask представляет связь между учеником и заданием
type UserTask struct {
	UserID     string    `json:"user_id"`     // ID пользователя (ученика)
	TaskID     string    `json:"task_id"`     // ID задания
	Completed  bool      `json:"completed"`   // Статус выполнения
	CompletedAt time.Time `json:"completed_at"` // Дата выполнения
}

// NewUserTask создает новую запись о выполнении задания
func NewUserTask(userID, taskID string) UserTask {
	return UserTask{
		UserID:     userID,
		TaskID:     taskID,
		Completed:  true,
		CompletedAt: time.Now(),
	}
}

// UserTaskStats содержит статистику выполнения заданий для пользователя
type UserTaskStats struct {
	UserID      string `json:"user_id"`
	CourseID    string `json:"course_id"`
	TotalTasks  int    `json:"total_tasks"`
	CompletedTasks int `json:"completed_tasks"`
	Progress    float64 `json:"progress"` // Процент выполнения (0-100)
}
