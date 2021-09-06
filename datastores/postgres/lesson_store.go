package postgres

import (
	"fmt"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	queryGetLesson          = `SELECT * FROM lessons WHERE id = $1`
	queryGetLessonsByCourse = `SELECT * FROM lessons WHERE course_id = $1`
	queryCreateLesson       = `INSERT INTO lessons VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	queryUpdateLesson       = `UPDATE lessons SET course_id = $1, title = $2, slug = $3, category = $4, sorting_order = $5 WHERE id = $6 RETURNING *`
	queryDeleteLesson       = `DELETE FROM lessons WHERE id = $1`
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

func (s *LessonStore) LessonsByCourse(courseID uuid.UUID) ([]courses.Lesson, error) {
	var ll []courses.Lesson
	if err := s.Select(&ll, queryGetLessonsByCourse, courseID); err != nil {
		return []courses.Lesson{}, fmt.Errorf("error getting lessons: %w", err)
	}
	return ll, nil
}

func (s *LessonStore) CreateLesson(l *courses.Lesson) error {
	if err := s.Get(l, queryCreateLesson,
		uuid.New(),
		l.CourseID,
		l.Title,
		l.Slug,
		l.Category,
		l.Order); err != nil {
		return fmt.Errorf("error creating lesson: %w", err)
	}
	return nil
}

func (s *LessonStore) UpdateLesson(l *courses.Lesson) error {
	var query = queryUpdateLesson
	if err := s.Get(l, query,
		l.CourseID,
		l.Title,
		l.Slug,
		l.Category,
		l.Order,
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
