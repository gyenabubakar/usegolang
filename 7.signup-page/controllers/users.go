package controllers

import (
	"fmt"
	"net/http"
	"signup_page/views"
)

func UsersController() *Users {
	return &Users{
		CreateUserView: views.NewView("bootstrap", "views/users/signup.gohtml"),
	}
}

func (u *Users) RenderSignupView(w http.ResponseWriter, r *http.Request) {
	u.CreateUserView.Render(w, nil, nil)
}

func (u *Users) HandlerUserCreation(w http.ResponseWriter, r *http.Request) {
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
	CreateUserView *views.View
}
