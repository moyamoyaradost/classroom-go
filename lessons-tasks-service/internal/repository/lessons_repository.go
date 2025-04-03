package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/models"
)

type LessonRepo struct {
	db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db: db}
}

func (r *LessonRepo) CreateLesson(ctx context.Context, lesson models.Lesson) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO lessons (id, course_id, description) VALUES ($1, $2, $3)",
		lesson.ID, lesson.CourseID, lesson.Description,
	)
	return err
}

func (r *LessonRepo) GetLessonByID(ctx context.Context, id string) (models.Lesson, error) {
	var lesson models.Lesson
	err := r.db.QueryRowContext(ctx,
		"SELECT id, course_id, description, created_at FROM lessons WHERE id = $1",
		id,
	).Scan(&lesson.ID, &lesson.CourseID, &lesson.Description, &lesson.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return lesson, errors.New("lesson not found")
	}
	return lesson, err
}

func (r *LessonRepo) GetLessonsByCourse(ctx context.Context, courseID string) ([]models.Lesson, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, course_id, description, created_at FROM lessons WHERE course_id = $1",
		courseID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []models.Lesson
	for rows.Next() {
		var lesson models.Lesson
		if err := rows.Scan(&lesson.ID, &lesson.CourseID, &lesson.Description, &lesson.CreatedAt); err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
