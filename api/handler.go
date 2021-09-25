package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
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
			r.Post("/", h.authRequest(courses.Create()))
			r.Put("/{id}", h.authRequest(courses.Update()))
			r.Delete("/{id}", h.authRequest(courses.Delete()))

			r.Get("/{courseID}/lessons/{lessonID}", lessons.Show())
			r.Get("/{courseID}/lessons", lessons.List())
			r.Post("/{courseID}/lessons", h.authRequest(lessons.Create()))
			r.Put("/{courseID}/lessons/{lessonID}", h.authRequest(lessons.Update()))
			r.Delete("/{courseID}/lessons/{lessonID}", h.authRequest(lessons.Delete()))
		})
		r.Route("/subjects", func(r chi.Router) {
			r.Get("/{id}", subjects.Show())
			r.Get("/", subjects.List())
			r.Post("/", h.authRequest(subjects.Create()))
			r.Put("/{id}", h.authRequest(subjects.Update()))
			r.Delete("/{id}", h.authRequest(subjects.Delete()))
		})
		r.Route("/jobs", func(r chi.Router) {
			r.Post("/seed", h.authRequest(jobs.Seed()))
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
