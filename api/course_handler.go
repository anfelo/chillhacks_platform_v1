package api

import (
	"net/http"

	"github.com/anfelo/chillhacks_platform/courses"
)

type CourseHandler struct {
	store courses.Store
}

func (h *CourseHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
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
