package postgres

import (
	"fmt"
	"time"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	queryGetLesson          = `SELECT * FROM lessons WHERE id = $1`
	queryGetLessons         = `SELECT * FROM lessons`
	queryGetLessonsByCourse = `
		SELECT * FROM lessons
		WHERE course_id = $1
		ORDER BY sorting_order ASC
	`
	queryCreateLesson       = `INSERT INTO lessons(id, course_id, title, content, slug, category, sorting_order, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *`
	queryUpdateLesson       = `UPDATE lessons SET course_id = $1, title = $2, content = $3, slug = $4, category = $5, sorting_order = $6, updated_at = $7 WHERE id = $8 RETURNING *`
	queryDeleteLesson       = `DELETE FROM lessons WHERE id = $1`
	queryCreateLessonsTable = `
		CREATE TABLE lessons (
			id UUID PRIMARY KEY,
			course_id UUID NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			slug TEXT NOT NULL UNIQUE,
			category TEXT NOT NULL,
			sorting_order INT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`
)

type LessonStore struct {
	*sqlx.DB
}

func (s *LessonStore) Lesson(id uuid.UUID) (courses.Lesson, error) {
	var l courses.Lesson
	if err := s.Get(&l, queryGetLesson, id); err != nil {
		return courses.Lesson{}, fmt.Errorf("error getting lesson: %w", err)
	}
	return l, nil
}

func (s *LessonStore) Lessons() ([]courses.Lesson, error) {
	var ll []courses.Lesson
	if err := s.Select(&ll, queryGetLessons); err != nil {
		return []courses.Lesson{}, fmt.Errorf("error getting lessons: %w", err)
	}
	return ll, nil
}

func (s *LessonStore) LessonsByCourse(courseID uuid.UUID) ([]courses.Lesson, error) {
	var ll []courses.Lesson
	if err := s.Select(&ll, queryGetLessonsByCourse, courseID); err != nil {
		return []courses.Lesson{}, fmt.Errorf("error getting lessons: %w", err)
	}
	return ll, nil
}

func (s *LessonStore) CreateLesson(l *courses.Lesson) error {
	l.ID = uuid.New()
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()
	if err := s.Get(l, queryCreateLesson,
		l.ID,
		l.CourseID,
		l.Title,
		l.Content,
		l.Slug,
		l.Category,
		l.Order,
		l.CreatedAt,
		l.UpdatedAt); err != nil {
		return fmt.Errorf("error creating lesson: %w", err)
	}
	return nil
}

func (s *LessonStore) UpdateLesson(l *courses.Lesson) error {
	l.UpdatedAt = time.Now()
	var query = queryUpdateLesson
	if err := s.Get(l, query,
		l.CourseID,
		l.Title,
		l.Content,
		l.Slug,
		l.Category,
		l.Order,
		l.UpdatedAt,
		l.ID); err != nil {
		return fmt.Errorf("error updating lesson: %w", err)
	}
	return nil
}

func (s *LessonStore) DeleteLesson(id uuid.UUID) error {
	if _, err := s.Exec(queryDeleteLesson, id); err != nil {
		return fmt.Errorf("error deleting lesson: %w", err)
	}
	return nil
}
