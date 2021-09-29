package api

import (
	"context"
	"net/http"

	"github.com/anfelo/chillhacks_platform/utils/errors"
	"github.com/anfelo/chillhacks_platform/utils/http_utils"
	"github.com/google/uuid"
)

func (h *Handler) withUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := h.sessions.Get(r.Context(), "user_id").(uuid.UUID)

		user, err := h.store.User(id)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) authRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := GetSessionData(h.sessions, r.Context())
		if s.LoggedIn {
			next.ServeHTTP(w, r)
			return
		}
		restErr := errors.NewUnauthorizedError("unauthorized request")
		http_utils.RespondJson(w, restErr.Status, restErr)
	})
}

func (h *Handler) adminRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := GetSessionData(h.sessions, r.Context())
		if s.User.Role == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		restErr := errors.NewUnauthorizedError("unauthorized request")
		http_utils.RespondJson(w, restErr.Status, restErr)
	})
}
