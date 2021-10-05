package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/anfelo/chillhacks_platform/utils/errors"
	"github.com/anfelo/chillhacks_platform/utils/http_utils"
)

func (h *Handler) withUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			next.ServeHTTP(w, r)
			return
		}
		username, tokenErr := VerifyJWT(r.Header["Authorization"][0])
		if tokenErr != nil {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println(username)

		user, err := h.store.UserByUsername(username)
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
		if r.Context().Value("user") != nil {
			user := r.Context().Value("user").(courses.User)
			if user.Username != "" {
				next.ServeHTTP(w, r)
				return
			}
		}
		restErr := errors.NewUnauthorizedError("unauthorized request")
		http_utils.RespondJson(w, restErr.Status, restErr)
	})
}

func (h *Handler) adminRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("user") != nil {
			user := r.Context().Value("user").(courses.User)
			if user.Role == "admin" {
				next.ServeHTTP(w, r)
				return
			}
		}
		restErr := errors.NewUnauthorizedError("unauthorized request")
		http_utils.RespondJson(w, restErr.Status, restErr)
	})
}
