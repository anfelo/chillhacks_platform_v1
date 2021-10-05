package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/anfelo/chillhacks_platform/courses"
	"github.com/anfelo/chillhacks_platform/utils/errors"
	"github.com/anfelo/chillhacks_platform/utils/http_utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	store    courses.Store
	sessions *scs.SessionManager
}

type AuthResponse struct {
	User  courses.User `json:"user"`
	Token string       `json:"token"`
}

func (h *UserHandler) CurrentUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("user") != nil {
			user := r.Context().Value("user").(courses.User)
			http_utils.RespondJson(w, http.StatusOK, user)
			return
		}
		http_utils.RespondJson(w, http.StatusOK, map[string]interface{}{"user": nil})
	}
}

func (h *UserHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		defer r.Body.Close()

		form := RegisterForm{
			UsernameTaken: false,
		}
		if err := json.Unmarshal(reqBody, &form); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		if _, err := h.store.UserByUsername(form.Username); err == nil {
			form.UsernameTaken = true
		}
		if !form.Validate() {
			restErr := errors.NewBadRequestError("invalid form values")
			restErr.Errors = form.Errors
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}

		password, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
		if err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
		}

		user := courses.User{
			ID:       uuid.New(),
			Username: form.Username,
			Password: string(password),
			Role:     form.Role,
		}
		if err := h.store.CreateUser(&user); err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		http_utils.RespondJson(w, http.StatusOK, user)
	}
}

func (h *UserHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}
		defer r.Body.Close()

		form := LoginForm{
			IncorrectCredentials: false,
		}
		if err := json.Unmarshal(reqBody, &form); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}

		user, err := h.store.UserByUsername(form.Username)
		if err != nil {
			form.IncorrectCredentials = true
		} else {
			compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
			form.IncorrectCredentials = compareErr != nil
		}
		if !form.Validate() {
			restErr := errors.NewBadRequestError("invalid form values")
			restErr.Errors = form.Errors
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}

		validToken, err := GenerateJWT(user.Username, user.ID.String())
		if err != nil {
			restErr := errors.NewInternatServerError("internal server error")
			http_utils.RespondJson(w, restErr.Status, restErr)
			return
		}

		authRes := AuthResponse{User: user, Token: validToken}
		http_utils.RespondJson(w, http.StatusOK, authRes)
	}
}

func (h *UserHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.sessions.Remove(r.Context(), "user")
		http_utils.RespondJson(w, http.StatusOK, map[string]string{"success": "true"})
	}
}
