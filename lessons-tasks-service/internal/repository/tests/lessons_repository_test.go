package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/internal/repository"
	"github.com/moyamoyaradost/classroom-go/lessons-tasks-service/models"
	"github.com/stretchr/testify/assert"
)

func TestLessonRepo_CreateLesson(t *testing.T) {
	// Настройка мок-базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Создание репозитория с мок-базой
	repo := repository.NewLessonRepo(db)

	// Подготовка тестовых данных
	lesson := models.Lesson{
		ID:          "test-id",
		CourseID:    "course-id",
		Description: "Test lesson",
	}

	// Настройка ожидаемого запроса
	mock.ExpectExec("INSERT INTO lessons").
		WithArgs(lesson.ID, lesson.CourseID, lesson.Description).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Выполнение тестируемого метода
	err = repo.CreateLesson(context.Background(), lesson)

	// Проверки
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLessonRepo_GetLessonByID(t *testing.T) {
	// Настройка мок-базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Создание репозитория с мок-базой
	repo := repository.NewLessonRepo(db)

	// Подготовка тестовых данных
	id := "test-id"
	courseID := "course-id"
	description := "Test lesson"
	createdAt := time.Now()

	// Настройка ожидаемого запроса с использованием регулярного выражения для SQL-запроса
	rows := sqlmock.NewRows([]string{"id", "course_id", "description", "created_at"}).
		AddRow(id, courseID, description, createdAt)

	mock.ExpectQuery("SELECT (.+) FROM lessons WHERE id = \\$1").
		WithArgs(id).
		WillReturnRows(rows)

	// Выполнение тестируемого метода
	lesson, err := repo.GetLessonByID(context.Background(), id)

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, id, lesson.ID)
	assert.Equal(t, courseID, lesson.CourseID)
	assert.Equal(t, description, lesson.Description)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLessonRepo_GetLessonsByCourse(t *testing.T) {
	// Настройка мок-базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Создание репозитория с мок-базой
	repo := repository.NewLessonRepo(db)

	// Подготовка тестовых данных
	courseID := "course-id"
	createdAt := time.Now()
	lessons := []models.Lesson{
		{ID: "lesson-1", CourseID: courseID, Description: "First lesson", CreatedAt: createdAt},
		{ID: "lesson-2", CourseID: courseID, Description: "Second lesson", CreatedAt: createdAt},
	}

	// Настройка ожидаемого запроса
	rows := sqlmock.NewRows([]string{"id", "course_id", "description", "created_at"})
	for _, l := range lessons {
		rows.AddRow(l.ID, l.CourseID, l.Description, l.CreatedAt)
	}

	mock.ExpectQuery("SELECT (.+) FROM lessons WHERE course_id = \\$1").
		WithArgs(courseID).
		WillReturnRows(rows)

	// Выполнение тестируемого метода
	result, err := repo.GetLessonsByCourse(context.Background(), courseID)

	// Проверки
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "lesson-1", result[0].ID)
	assert.Equal(t, "lesson-2", result[1].ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLessonRepo_DeleteLesson(t *testing.T) {
	// Настройка мок-базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Создание репозитория с мок-базой
	repo := repository.NewLessonRepo(db)

	// Подготовка тестовых данных
	lessonID := "lesson-1"

	// Настройка ожидаемого запроса
	mock.ExpectExec("DELETE FROM lessons WHERE id = \\$1").
		WithArgs(lessonID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Выполнение тестируемого метода
	err = repo.DeleteLesson(context.Background(), lessonID)

	// Проверки
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
