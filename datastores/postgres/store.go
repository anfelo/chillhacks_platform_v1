package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		DB:           db,
		CourseStore:  &CourseStore{DB: db},
		LessonStore:  &LessonStore{DB: db},
		SubjectStore: &SubjectStore{DB: db},
		UserStore:    &UserStore{DB: db},
	}, nil
}

type Store struct {
	*sqlx.DB

	*CourseStore
	*LessonStore
	*SubjectStore
	*UserStore
}

func (s *Store) RunMigration(query string) error {
	if _, err := s.Exec(query); err != nil {
		return fmt.Errorf("error running migration: %w", err)
	}
	return nil
}
