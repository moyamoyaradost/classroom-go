// models/task_answer.go
package models

import "time"

type TaskAnswer struct {
	ID             string    `json:"id"`
	TaskID         string    `json:"task_id"`
	UserID         string    `json:"user_id"`
	AnswerText     string    `json:"answer_text"`
	AttachmentURLs []string  `json:"attachment_urls"`
	Status         string    `json:"status"` // "pending", "approved", "rejected"
	CreatedAt      time.Time `json:"created_at"`
}
