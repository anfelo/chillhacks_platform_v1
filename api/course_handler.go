package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/anfelo/chillhacks_platform/utils/http_utils"
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

		http_utils.RespondJson(w, http.StatusOK, c)
	}
}

func (h *CourseHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		suIdStr := strings.TrimSpace(r.URL.Query().Get("subject"))
		var ccRes courses.CoursesResponse
		if suIdStr != "" {
			suID, err := uuid.Parse(suIdStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			cc, err := h.store.CoursesBySubject(suID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			ccRes = courses.CoursesResponse{
				Count:   len(cc),
				Results: cc,
			}
		} else {
			cc, err := h.store.Courses()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			ccRes = courses.CoursesResponse{
				Count:   len(cc),
				Results: cc,
			}
		}

		http_utils.RespondJson(w, http.StatusOK, ccRes)
	}
}

func (h *CourseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var c courses.Course
		if err := json.Unmarshal(reqBody, &c); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.CreateCourse(&c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http_utils.RespondJson(w, http.StatusCreated, c)
	}
}

func (h *CourseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var c courses.Course
		if err := json.Unmarshal(reqBody, &c); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.UpdateCourse(&c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http_utils.RespondJson(w, http.StatusOK, c)
	}
}

func (h *CourseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		err = h.store.DeleteCourse(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http_utils.RespondJson(w, http.StatusOK, map[string]string{"success": "true"})
	}
}
