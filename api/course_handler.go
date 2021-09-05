package api

import (
	"net/http"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type CourseHandler struct {
	store courses.Store
}

func (h *CourseHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		c, err := h.store.Course(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		serializer := &courses.CourseSerializer{}
		resBody, err := serializer.Encode(&c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		setupResponse(w, "application/json", resBody, http.StatusOK)
	}
}

func (h *CourseHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *CourseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *CourseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *CourseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
