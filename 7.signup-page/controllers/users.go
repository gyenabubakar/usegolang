package controllers

import (
	"fmt"
	"net/http"
	"signup_page/views"
)

func UsersController() *Users {
	return &Users{
		SignupView: views.NewView("bootstrap", "users/signup"),
	}
}

func (u *Users) RenderSignupView(w http.ResponseWriter, r *http.Request) {
	u.SignupView.Render(w, nil, nil)
}

func (u *Users) HandleUserCreation(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html")

	var form SignupForm
	if err := ParseForm(r, &form); err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintln(w, form)
}



type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type Users struct {
	SignupView *views.View
}
