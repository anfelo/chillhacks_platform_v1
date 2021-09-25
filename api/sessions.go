package api

import (
	"context"
	"database/sql"
	"encoding/gob"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/google/uuid"
)

func init() {
	gob.Register(uuid.UUID{})
}

func NewSessionManager(dataSourceName string) (*scs.SessionManager, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	sessions := scs.New()
	sessions.Store = postgresstore.New(db)

	return sessions, nil
}

type SessionData struct {
	User     courses.User `json:"user"`
	LoggedIn bool         `json:"-"`
}

func GetSessionData(session *scs.SessionManager, ctx context.Context) SessionData {
	var data SessionData
	data.User, data.LoggedIn = ctx.Value("user").(courses.User)
	return data
}
