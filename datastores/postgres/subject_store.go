package postgres

import (
	"fmt"
	"time"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	queryGetSubject    = `SELECT * FROM subjects WHERE id = $1`
	queryGetSubjects   = `SELECT * FROM subjects`
	queryCreateSubject = `INSERT INTO subjects(id, title, slug, img_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	queryUpdateSubject = `UPDATE subjects SET title = $1, slug = $2, img_url = $3, updated_at = $4 WHERE id = $5 RETURNING *`
	queryDeleteSubject = `DELETE FROM subjects WHERE id = $1`
)

type SubjectStore struct {
	*sqlx.DB
}

func (s *SubjectStore) Subject(id uuid.UUID) (courses.Subject, error) {
	var su courses.Subject
	if err := s.Get(&su, queryGetSubject, id); err != nil {
		return courses.Subject{}, fmt.Errorf("error getting course: %w", err)
	}
	return su, nil
}

func (s *SubjectStore) Subjects() ([]courses.Subject, error) {
	var ss []courses.Subject
	if err := s.Select(&ss, queryGetSubjects); err != nil {
		return []courses.Subject{}, fmt.Errorf("error getting subjects: %w", err)
	}
	return ss, nil
}

func (s *SubjectStore) CreateSubject(su *courses.Subject) error {
	su.ID = uuid.New()
	su.CreatedAt = time.Now()
	su.UpdatedAt = time.Now()
	if err := s.Get(su, queryCreateSubject,
		su.ID,
		su.Title,
		su.Slug,
		su.ImgURL,
		su.CreatedAt,
		su.UpdatedAt); err != nil {
		return fmt.Errorf("error creating subject: %w", err)
	}
	return nil
}

func (s *SubjectStore) UpdateSubject(su *courses.Subject) error {
	su.UpdatedAt = time.Now()
	if err := s.Get(su, queryUpdateSubject,
		su.Title,
		su.Slug,
		su.ImgURL,
		su.UpdatedAt,
		su.ID); err != nil {
		return fmt.Errorf("error updating subject: %w", err)
	}
	return nil
}

func (s *SubjectStore) DeleteSubject(id uuid.UUID) error {
	if _, err := s.Exec(queryDeleteSubject, id); err != nil {
		return fmt.Errorf("error deleting subject: %w", err)
	}
	return nil
}
