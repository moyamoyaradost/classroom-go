// internal/repository/task_answer_repository.go
package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/models"
)

type TaskAnswerRepo struct {
	db *sql.DB
}

func NewTaskAnswerRepo(db *sql.DB) *TaskAnswerRepo {
	return &TaskAnswerRepo{db: db}
}

func (r *TaskAnswerRepo) SaveTaskAnswer(ctx context.Context, answer models.TaskAnswer) error {
	// Сериализуем массив URL вложений в JSON
	attachmentsJSON, err := json.Marshal(answer.AttachmentURLs)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx,
		`INSERT INTO task_answers (id, task_id, user_id, answer_text, attachment_urls, status) 
         VALUES ($1, $2, $3, $4, $5, $6)
         ON CONFLICT (task_id, user_id) DO UPDATE 
         SET answer_text = $4, attachment_urls = $5, status = $6`,
		answer.ID, answer.TaskID, answer.UserID, answer.AnswerText, attachmentsJSON, answer.Status,
	)
	return err
}

func (r *TaskAnswerRepo) GetTaskAnswersByUser(ctx context.Context, taskID, userID string) ([]models.TaskAnswer, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, task_id, user_id, answer_text, attachment_urls, status, created_at 
         FROM task_answers 
         WHERE task_id = $1 AND user_id = $2`,
		taskID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTaskAnswers(rows)
}

func (r *TaskAnswerRepo) GetAllTaskAnswers(ctx context.Context, taskID string) ([]models.TaskAnswer, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, task_id, user_id, answer_text, attachment_urls, status, created_at 
         FROM task_answers 
         WHERE task_id = $1`,
		taskID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanTaskAnswers(rows)
}

// Вспомогательная функция для сканирования результатов запроса
func scanTaskAnswers(rows *sql.Rows) ([]models.TaskAnswer, error) {
	var answers []models.TaskAnswer

	for rows.Next() {
		var answer models.TaskAnswer
		var attachmentsJSON []byte

		if err := rows.Scan(
			&answer.ID,
			&answer.TaskID,
			&answer.UserID,
			&answer.AnswerText,
			&attachmentsJSON,
			&answer.Status,
			&answer.CreatedAt,
		); err != nil {
			return nil, err
		}

		// Десериализуем JSON в массив URL
		if err := json.Unmarshal(attachmentsJSON, &answer.AttachmentURLs); err != nil {
			return nil, err
		}

		answers = append(answers, answer)
	}

	return answers, nil
}
