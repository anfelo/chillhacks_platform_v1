package postgres

import (
	"fmt"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	queryGetCourse  = `SELECT * FROM courses WHERE id = $1`
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
	queryCreateCourse = `INSERT INTO courses VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	queryUpdateCourse = `UPDATE courses SET subject_id = $1, title = $2, description = $3, slug = $4, img_url = $5 WHERE id = $6 RETURNING *`
	queryDeleteCourse = `DELETE FROM courses WHERE id = $1`
)

type CourseStore struct {
	*sqlx.DB
}

func (s *CourseStore) Course(id uuid.UUID) (courses.Course, error) {
	var t courses.Course
	if err := s.Get(&t, queryGetCourse, id); err != nil {
		return courses.Course{}, fmt.Errorf("error getting course: %w", err)
	}
	return t, nil
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
	if err := s.Get(c, queryCreateCourse,
		c.ID,
		c.SubjectID,
		c.Title,
		c.Description,
		c.Slug,
		c.ImgURL); err != nil {
		return fmt.Errorf("error creating course: %w", err)
	}
	return nil
}

func (s *CourseStore) UpdateCourse(c *courses.Course) error {
	if err := s.Get(c, queryUpdateCourse,
		c.SubjectID,
		c.Title,
		c.Description,
		c.Slug,
		c.ImgURL,
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
