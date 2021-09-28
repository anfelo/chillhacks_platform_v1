package courses

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"-" db:"password"`
	Role     string    `json:"role" db:"role"`
}

type UserStore interface {
	CreateSessionsTable() error
	CreateUsersTable() error
	User(id uuid.UUID) (User, error)
	UserByUsername(username string) (User, error)
	CreateUser(t *User) error
	UpdateUser(t *User) error
	DeleteUser(id uuid.UUID) error
}
