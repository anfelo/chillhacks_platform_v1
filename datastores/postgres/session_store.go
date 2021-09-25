package postgres

import (
	"fmt"
)

const queryCreateSessionsTable = `
		CREATE TABLE sessions (
			token TEXT PRIMARY KEY,
			data BYTEA NOT NULL,
			expiry TIMESTAMPTZ NOT NULL
		);
		
		CREATE INDEX sessions_expiry_idx ON sessions (expiry)
	`

func (s *UserStore) CreateSessionsTable() error {
	if _, err := s.Exec(queryCreateSessionsTable); err != nil {
		return fmt.Errorf("error creating sessions table: %w", err)
	}
	return nil
}
