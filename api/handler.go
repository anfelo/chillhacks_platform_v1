package api

import (
	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHandler(store courses.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	courses := CourseHandler{store: store}
	lessons := LessonHandler{store: store}
	subjects := SubjectHandler{store: store}
	seeder := SeederHandler{store: store}

	h.Use(middleware.Logger)

	h.Route("/courses", func(r chi.Router) {
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
	h.Route("/subjects", func(r chi.Router) {
		r.Get("/{id}", subjects.Show())
		r.Get("/", subjects.List())
		r.Post("/", subjects.Create())
		r.Put("/{id}", subjects.Update())
		r.Delete("/{id}", subjects.Delete())
	})
	h.Post("/seed", seeder.Seed())

	return h
}

type Handler struct {
	*chi.Mux

	store courses.Store
}
