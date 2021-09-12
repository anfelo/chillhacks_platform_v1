package api

import (
	"context"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

func NewHandler(
	store courses.Store,
	sessions *scs.SessionManager) *Handler {
	h := &Handler{
		Mux:      chi.NewMux(),
		store:    store,
		sessions: sessions,
	}

	courses := CourseHandler{store: store, sessions: sessions}
	lessons := LessonHandler{store: store, sessions: sessions}
	subjects := SubjectHandler{store: store, sessions: sessions}
	jobs := JobHandler{store: store, sessions: sessions}
	users := UserHandler{store: store, sessions: sessions}

	h.Use(middleware.Logger)
	h.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*", "*"},
	}))
	h.Use(sessions.LoadAndSave)
	h.Use(h.withUser)

	h.Route("/api", func(r chi.Router) {
		r.Route("/courses", func(r chi.Router) {
			r.Get("/{id}", courses.Show())
			r.Get("/", courses.List())
			r.Post("/", courses.Create())
			r.Put("/{id}", courses.Update())
			r.Delete("/{id}", courses.Delete())

			r.Get("/{courseID}/lessons/{lessonID}", lessons.Show())
			r.Get("/{courseID}/lessons", lessons.List())
			r.Post("/{courseID}/lessons", lessons.Create())
			r.Put("/{courseID}/lessons/{lessonID}", lessons.Update())
			r.Delete("/{courseID}/lessons/{lessonID}", lessons.Delete())
		})
		r.Route("/subjects", func(r chi.Router) {
			r.Get("/{id}", subjects.Show())
			r.Get("/", subjects.List())
			r.Post("/", subjects.Create())
			r.Put("/{id}", subjects.Update())
			r.Delete("/{id}", subjects.Delete())
		})
		r.Route("/jobs", func(r chi.Router) {
			r.Post("/seed", jobs.Seed())
			r.Post("/create-tables", jobs.CreateTables())
		})
	})

	h.Route("/auth", func(r chi.Router) {
		r.Get("/currentuser", users.CurrentUser())
		r.Post("/register", users.RegisterSubmit())
		r.Post("/login", users.LoginSubmit())
		r.Get("/logout", users.Logout())
	})

	return h
}

type Handler struct {
	*chi.Mux

	store    courses.Store
	sessions *scs.SessionManager
}

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
