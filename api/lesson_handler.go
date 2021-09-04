package api

import (
	"net/http"

	"github.com/anfelo/chillhacks_platform/courses"
)

type LessonHandler struct {
	store courses.Store
}

func (h *LessonHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *LessonHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *LessonHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *LessonHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *LessonHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
