package postgres

import (
	"fmt"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	queryGetUser           = `SELECT * FROM users WHERE id = $1`
	queryGetUserByUsername = `SELECT * FROM users WHERE username = $1`
	queryGetUsers          = `SELECT * FROM users`
	queryCreateUser        = `INSERT INTO users VALUES ($1, $2, $3, $4) RETURNING *`
	queryUpdateUser        = `UPDATE users SET username = $1, password = $2, role = $3 WHERE id = $4 RETURNING *`
	queryDeleteUser        = `DELETE FROM users WHERE id = $1`
	queryCreateUsersTable  = `
		CREATE TABLE users (
			id UUID PRIMARY KEY,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) CreateUsersTable() error {
	if _, err := s.Exec(queryCreateUsersTable); err != nil {
		return fmt.Errorf("error creating users table: %w", err)
	}
	return nil
}

func (s *UserStore) User(id uuid.UUID) (courses.User, error) {
	var u courses.User
	if err := s.Get(&u, queryGetUser, id); err != nil {
		return courses.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return u, nil
}

func (s *UserStore) UserByUsername(username string) (courses.User, error) {
	var u courses.User
	if err := s.Get(&u, queryGetUserByUsername, username); err != nil {
		return courses.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return u, nil
}

func (s *UserStore) Users() ([]courses.User, error) {
	var uu []courses.User
	if err := s.Select(&uu, queryGetUsers); err != nil {
		return []courses.User{}, fmt.Errorf("error getting users: %w", err)
	}
	return uu, nil
}

func (s *UserStore) CreateUser(u *courses.User) error {
	if err := s.Get(u, queryCreateUser,
		u.ID,
		u.Username,
		u.Password,
		u.Role); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (s *UserStore) UpdateUser(u *courses.User) error {
	if err := s.Get(u, queryUpdateUser,
		u.Username,
		u.Password,
		u.Role,
		u.ID); err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (s *UserStore) DeleteUser(id uuid.UUID) error {
	if _, err := s.Exec(queryDeleteUser, id); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
