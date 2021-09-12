package api

import "encoding/gob"

func init() {
	gob.Register(RegisterForm{})
	gob.Register(LoginForm{})
	gob.Register(FormErrors{})
}

type FormErrors map[string]string

type RegisterForm struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	UsernameTaken bool   `json:"-"`

	Errors FormErrors `json:"errors"`
}

func (f *RegisterForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Username == "" {
		f.Errors["username"] = "Please enter a username."
	} else if f.UsernameTaken {
		f.Errors["username"] = "This username is already taken."
	}

	if f.Password == "" {
		f.Errors["password"] = "Please enter a password."
	} else if len(f.Password) < 8 {
		f.Errors["password"] = "Your password must be at least 8 characters long."
	}

	return len(f.Errors) == 0
}

type LoginForm struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	IncorrectCredentials bool   `json:"-"`

	Errors FormErrors `json:"errors"`
}

func (f *LoginForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Username == "" {
		f.Errors["username"] = "Please enter a username."
	} else if f.IncorrectCredentials {
		f.Errors["username"] = "Username or password is incorrect."
	}

	if f.Password == "" {
		f.Errors["password"] = "Please enter a password."
	}

	return len(f.Errors) == 0
}
