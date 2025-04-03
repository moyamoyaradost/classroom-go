package repository

import (
	"context"
	"database/sql"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/models"
)

type UserTaskRepo struct {
	db *sql.DB
}

func NewUserTaskRepo(db *sql.DB) *UserTaskRepo {
	return &UserTaskRepo{db: db}
}

func (r *UserTaskRepo) CompleteTask(ctx context.Context, userID, taskID string) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO user_tasks (user_id, task_id) VALUES ($1, $2) ON CONFLICT DO NOTHING",
		userID, taskID,
	)
	return err
}

func (r *UserTaskRepo) GetCompletedTasks(ctx context.Context, userID string) ([]models.Task, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT t.id, t.course_id, t.description, t.created_at 
		FROM tasks t
		JOIN user_tasks ut ON t.id = ut.task_id
		WHERE ut.user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.CourseID, &task.Description, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
