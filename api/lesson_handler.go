package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/anfelo/chillhacks_platform/utils/errors"
	"github.com/anfelo/chillhacks_platform/utils/http_utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type LessonHandler struct {
	store    courses.Store
	sessions *scs.SessionManager
}

func (h *LessonHandler) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lIDStr := chi.URLParam(r, "lessonID")
		lID, err := uuid.Parse(lIDStr)
		if err != nil {
			restErr := errors.NewBadRequestError("invalid lesson id")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}

		l, err := h.store.Lesson(lID)
		if err != nil {
			restErr := errors.NewNotFoundError("lesson not found")
			http_utils.RespondJson(w, restErr.Status, restErr)
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
			restErr := errors.NewBadRequestError("invalid course id")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}

		ll, err := h.store.LessonsByCourse(cID)
		if err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
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
			restErr := errors.NewBadRequestError("invalid course id")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		defer r.Body.Close()

		var l courses.Lesson
		if err := json.Unmarshal(reqBody, &l); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		l.CourseID = cID

		if err := h.store.CreateLesson(&l); err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
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
			restErr := errors.NewBadRequestError("invalid course id")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		defer r.Body.Close()

		var l courses.Lesson
		if err := json.Unmarshal(reqBody, &l); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		l.CourseID = cID

		if err := h.store.UpdateLesson(&l); err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
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
			restErr := errors.NewBadRequestError("invalid lesson id")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}

		err = h.store.DeleteLesson(lID)
		if err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}

		http_utils.RespondJson(w, http.StatusOK, map[string]string{"success": "true"})
	}
}
