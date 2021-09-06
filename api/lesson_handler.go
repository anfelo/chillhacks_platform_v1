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

type LessonHandler struct {
	store courses.Store
}

func (h *LessonHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lIDStr := chi.URLParam(r, "lessonID")
		lID, err := uuid.Parse(lIDStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		l, err := h.store.Lesson(lID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http_utils.RespondJson(w, http.StatusOK, l)
	}
}

func (h *LessonHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cIDStr := chi.URLParam(r, "courseID")
		cID, err := uuid.Parse(cIDStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		ll, err := h.store.LessonsByCourse(cID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		llRes := courses.LessonsResponse{
			Count:   len(ll),
			Results: ll,
		}
		http_utils.RespondJson(w, http.StatusOK, llRes)
	}
}

func (h *LessonHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cIDStr := chi.URLParam(r, "courseID")
		cID, err := uuid.Parse(cIDStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var l courses.Lesson
		if err := json.Unmarshal(reqBody, &l); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		l.CourseID = cID

		if err := h.store.CreateLesson(&l); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http_utils.RespondJson(w, http.StatusCreated, l)
	}
}

func (h *LessonHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cIDStr := chi.URLParam(r, "courseID")
		cID, err := uuid.Parse(cIDStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var l courses.Lesson
		if err := json.Unmarshal(reqBody, &l); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		l.CourseID = cID

		if err := h.store.UpdateLesson(&l); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http_utils.RespondJson(w, http.StatusOK, l)
	}
}

func (h *LessonHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lIDStr := chi.URLParam(r, "lessonID")
		lID, err := uuid.Parse(lIDStr)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		err = h.store.DeleteLesson(lID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http_utils.RespondJson(w, http.StatusOK, map[string]string{"success": "true"})
	}
}
