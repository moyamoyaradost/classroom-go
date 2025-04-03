package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/models"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) CreateTask(ctx context.Context, task models.Task) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO tasks (id, course_id, description) VALUES ($1, $2, $3)",
		task.ID, task.CourseID, task.Description,
	)
	return err
}

func (r *TaskRepo) GetTaskByID(ctx context.Context, id string) (models.Task, error) {
	var task models.Task
	err := r.db.QueryRowContext(ctx,
		"SELECT id, course_id, description, created_at FROM tasks WHERE id = $1",
		id,
	).Scan(&task.ID, &task.CourseID, &task.Description, &task.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return task, errors.New("task not found")
	}
	return task, err
}
func (r *TaskRepo) GetTasksByCourse(ctx context.Context, courseID string) ([]models.Task, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, course_id, description, created_at FROM tasks WHERE course_id = $1",
		courseID,
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
