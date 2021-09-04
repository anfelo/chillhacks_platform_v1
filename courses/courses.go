package courses

import "github.com/google/uuid"

type Course struct {
	ID           uuid.UUID `db:"id"`
	SubjectID    uuid.UUID `db:"subject_id"`
	Title        string    `db:"title"`
	Description  string    `db:"description"`
	Slug         string    `db:"slug"`
	LessonsCount int       `db:"lessons_count"`
	ImgURL       string    `db:"img_url"`
}

type Lesson struct {
	ID       uuid.UUID `db:"id"`
	CourseID uuid.UUID `db:"course_id"`
	Title    string    `db:"title"`
	Slug     string    `db:"slug"`
	Category string    `db:"category"`
	Order    int       `db:"order"`
}

type Subject struct {
	ID     uuid.UUID `db:"id"`
	Title  string    `db:"title"`
	ImgURL string    `db:"img_url"`
}

type CourseStore interface {
	Course(id uuid.UUID) (Course, error)
	Courses() ([]Course, error)
	CoursesBySubject(subjectID uuid.UUID) ([]Course, error)
	CreateCourse(c *Course) error
	UpdateCourse(c *Course) error
	DeleteCourse(id uuid.UUID) error
}

type LessonStore interface {
	Lesson(id uuid.UUID) (Lesson, error)
	LessonsByCourse(courseID uuid.UUID) ([]Lesson, error)
	CreateLesson(l *Lesson) error
	UpdateLesson(l *Lesson) error
	DeleteLesson(id uuid.UUID) error
}

type SubjectStore interface {
	Subject(id uuid.UUID) (Subject, error)
	Subjects() ([]Subject, error)
	CreateSubject(s *Subject) error
	UpdateSubject(s *Subject) error
	DeleteSubject(id uuid.UUID) error
}

type Store interface {
	CourseStore
	LessonStore
	SubjectStore
}
