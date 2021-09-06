package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/anfelo/chillhacks_platform/utils/http_utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type SubjectHandler struct {
	store courses.Store
}

func (h *SubjectHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		s, err := h.store.Subject(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http_utils.RespondJson(w, http.StatusOK, s)
	}
}

func (h *SubjectHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss, err := h.store.Subjects()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ssRes := courses.SubjectsResponse{
			Count:   len(ss),
			Results: ss,
		}
		http_utils.RespondJson(w, http.StatusOK, ssRes)
	}
}

func (h *SubjectHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var s courses.Subject
		if err := json.Unmarshal(reqBody, &s); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.CreateSubject(&s); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http_utils.RespondJson(w, http.StatusCreated, s)
	}
}

func (h *SubjectHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var s courses.Subject
		if err := json.Unmarshal(reqBody, &s); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.UpdateSubject(&s); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http_utils.RespondJson(w, http.StatusOK, s)
	}
}

func (h *SubjectHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		err = h.store.DeleteSubject(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http_utils.RespondJson(w, http.StatusOK, map[string]string{"success": "true"})
	}
}
