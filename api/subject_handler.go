package api

import (
	"net/http"

	"github.com/anfelo/chillhacks_platform/courses"
)

type SubjectHandler struct {
	store courses.Store
}

func (h *SubjectHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *SubjectHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *SubjectHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *SubjectHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *SubjectHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
