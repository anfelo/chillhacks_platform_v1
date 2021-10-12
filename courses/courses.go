package courses

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID           uuid.UUID `json:"id" db:"id"`
	SubjectID    uuid.UUID `json:"subject_id" db:"subject_id"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	Slug         string    `json:"slug" db:"slug"`
	LessonsCount int       `json:"lessons_count" db:"lessons_count"`
	ImgURL       string    `json:"img_url" db:"img_url"`
	CreatedAt    time.Time `json:"created" db:"created_at"`
	UpdatedAt    time.Time `json:"updated" db:"updated_at"`
}

type CoursesResponse struct {
	Count   int      `json:"count"`
	Results []Course `json:"results"`
}

type Lesson struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CourseID  uuid.UUID `json:"course_id" db:"course_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Slug      string    `json:"slug" db:"slug"`
	Category  string    `json:"category" db:"category"`
	Order     int       `json:"sorting_order" db:"sorting_order"`
	CreatedAt time.Time `json:"created" db:"created_at"`
	UpdatedAt time.Time `json:"updated" db:"updated_at"`
}

type LessonsResponse struct {
	Count   int      `json:"count"`
	Results []Lesson `json:"results"`
}

type Subject struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	ImgURL    string    `json:"img_url" db:"img_url"`
	Slug      string    `json:"slug" db:"slug"`
	CreatedAt time.Time `json:"created" db:"created_at"`
	UpdatedAt time.Time `json:"updated" db:"updated_at"`
}

type SubjectsResponse struct {
	Count   int       `json:"count"`
	Results []Subject `json:"results"`
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
	Lessons() ([]Lesson, error)
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
	UserStore

	RunMigration(query string) error
}
