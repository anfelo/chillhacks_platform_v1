package postgres

import (
	"fmt"
	"time"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	queryGetCourse = `
		SELECT
			courses.*,
			COUNT(lessons.*) AS lessons_count
		FROM courses
		LEFT JOIN lessons ON lessons.course_id = courses.id
		WHERE courses.id = $1
		GROUP BY courses.id
	`
	queryGetCourses = `
		SELECT
			courses.*,
			COUNT(lessons.*) AS lessons_count
		FROM courses
		LEFT JOIN lessons ON lessons.course_id = courses.id
		GROUP BY courses.id
	`
	queryGetCoursesBySubjects = `
		SELECT
			courses.*,
			COUNT(lessons.*) AS lessons_count
		FROM courses
		LEFT JOIN lessons ON lessons.course_id = courses.id
		WHERE subject_id = $1
		GROUP BY courses.id
	`
	queryCreateCourse       = `INSERT INTO courses(id, subject_id, title, description, slug, img_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`
	queryUpdateCourse       = `UPDATE courses SET subject_id = $1, title = $2, description = $3, slug = $4, img_url = $5, updated_at = $6 WHERE id = $7 RETURNING *`
	queryDeleteCourse       = `DELETE FROM courses WHERE id = $1`
	queryCreateCoursesTable = `
		CREATE TABLE courses (
			id UUID PRIMARY KEY,
			subject_id UUID NOT NULL REFERENCES subjects (id) ON DELETE CASCADE,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			img_url TEXT NOT NULL,
			slug TEXT NOT NULL UNIQUE,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`
)

type CourseStore struct {
	*sqlx.DB
}

func (s *CourseStore) CreateCoursesTable() error {
	if _, err := s.Exec(queryCreateCoursesTable); err != nil {
		return fmt.Errorf("error creating courses table: %w", err)
	}
	return nil
}

func (s *CourseStore) Course(id uuid.UUID) (courses.Course, error) {
	var c courses.Course
	if err := s.Get(&c, queryGetCourse, id); err != nil {
		return courses.Course{}, fmt.Errorf("error getting course: %w", err)
	}
	return c, nil
}

func (s *CourseStore) Courses() ([]courses.Course, error) {
	var cc []courses.Course
	if err := s.Select(&cc, queryGetCourses); err != nil {
		return []courses.Course{}, fmt.Errorf("error getting courses: %w", err)
	}
	return cc, nil
}

func (s *CourseStore) CoursesBySubject(subjectID uuid.UUID) ([]courses.Course, error) {
	var cc []courses.Course
	if err := s.Select(&cc, queryGetCoursesBySubjects, subjectID); err != nil {
		return []courses.Course{}, fmt.Errorf("error getting courses: %w", err)
	}
	return cc, nil
}

func (s *CourseStore) CreateCourse(c *courses.Course) error {
	c.ID = uuid.New()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	if err := s.Get(c, queryCreateCourse,
		c.ID,
		c.SubjectID,
		c.Title,
		c.Description,
		c.Slug,
		c.ImgURL,
		c.CreatedAt,
		c.UpdatedAt); err != nil {
		return fmt.Errorf("error creating course: %w", err)
	}
	return nil
}

func (s *CourseStore) UpdateCourse(c *courses.Course) error {
	c.UpdatedAt = time.Now()
	if err := s.Get(c, queryUpdateCourse,
		c.SubjectID,
		c.Title,
		c.Description,
		c.Slug,
		c.ImgURL,
		c.UpdatedAt,
		c.ID); err != nil {
		return fmt.Errorf("error updating course: %w", err)
	}
	return nil
}

func (s *CourseStore) DeleteCourse(id uuid.UUID) error {
	if _, err := s.Exec(queryDeleteCourse, id); err != nil {
		return fmt.Errorf("error deleting course: %w", err)
	}
	return nil
}
