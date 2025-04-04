package models

import (
	"errors"
	"time"
	
	v1 "github.com/moyamoyaradost/classroom-go/lessons-tasks-service/api/v1"
)

type TaskAnswer struct {
	ID             string    `json:"id"`
	TaskID         string    `json:"task_id"`
	UserID         string    `json:"user_id"`
	AnswerText     string    `json:"answer_text"`
	AttachmentURLs []string  `json:"attachment_urls"`
	Status         string    `json:"status"` // "pending", "approved", "rejected"
	CreatedAt      time.Time `json:"created_at"`
}

// Конвертация из protobuf enum в строку
func StatusFromProto(status v1.AnswerStatus) string {
	return status.String()
}

// Конвертация строки в protobuf enum с валидацией
func StatusToProto(status string) (v1.AnswerStatus, error) {
	switch status {
	case v1.AnswerStatus_PENDING.String():
		return v1.AnswerStatus_PENDING, nil
	case v1.AnswerStatus_APPROVED.String():
		return v1.AnswerStatus_APPROVED, nil
	case v1.AnswerStatus_REJECTED.String():
		return v1.AnswerStatus_REJECTED, nil
	default:
		return v1.AnswerStatus_PENDING, errors.New("invalid status value")
	}
}
