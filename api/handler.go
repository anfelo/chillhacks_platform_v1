package api

import (
	"log"
	"net/http"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
)

func NewHandler(store courses.Store, csrfKey []byte) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	courses := CourseHandler{store: store}
	lessons := LessonHandler{store: store}
	subjects := SubjectHandler{store: store}

	h.Use(middleware.Logger)
	h.Use(csrf.Protect(csrfKey, csrf.Secure(false)))

	h.Route("/courses", func(r chi.Router) {
		r.Get("/{id}", courses.Show())
		r.Get("/", courses.List())
		r.Post("/", courses.Create())
		r.Put("/{id}", courses.Update())
		r.Delete("/{id}", courses.Delete())
	})
	h.Route("/lessons", func(r chi.Router) {
		r.Get("/{id}", lessons.Show())
		r.Get("/", lessons.List())
		r.Post("/", lessons.Create())
		r.Put("/{id}", lessons.Update())
		r.Delete("/{id}", lessons.Delete())
	})
	h.Route("/subjects", func(r chi.Router) {
		r.Get("/{id}", subjects.Show())
		r.Get("/", subjects.List())
		r.Post("/", subjects.Create())
		r.Put("/{id}", subjects.Update())
		r.Delete("/{id}", subjects.Delete())
	})

	return h
}

type Handler struct {
	*chi.Mux

	store courses.Store
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
