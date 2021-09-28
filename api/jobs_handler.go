package api

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/anfelo/chillhacks_platform/utils/http_utils"
)

type JobHandler struct {
	store    courses.Store
	sessions *scs.SessionManager
}

func (h *JobHandler) CreateTables() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errs := make(map[string]interface{})
		err := h.store.CreateSubjectsTable()
		if err != nil {
			errs["subjects"] = "table already created"
		} else {
			errs["subjects"] = "table created"
		}
		err = h.store.CreateCoursesTable()
		if err != nil {
			errs["courses"] = "table already created"
		} else {
			errs["courses"] = "table created"
		}
		err = h.store.CreateLessonsTable()
		if err != nil {
			errs["lessons"] = "table already created"
		} else {
			errs["lessons"] = "table created"
		}
		err = h.store.CreateUsersTable()
		if err != nil {
			errs["users"] = "table already created"
		} else {
			errs["users"] = "table created"
		}
		err = h.store.CreateSessionsTable()
		if err != nil {
			errs["sessions"] = "table already created"
		} else {
			errs["sessions"] = "table created"
		}
		http_utils.RespondJson(w, http.StatusCreated, errs)
	}
}

func (h *JobHandler) RunMigrations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Run migrations
	}
}
