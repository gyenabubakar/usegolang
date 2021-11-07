package controllers

import (
	"first-views/views"
	"net/http"
)

func UsersController() *Users {
	return &Users{
		CreateUserView: views.NewView("bootstrap", "views/users/create-user.gohtml"),
	}
}

func (u *Users) RenderCreationView(w http.ResponseWriter, r *http.Request) {
	u.CreateUserView.Render(w, nil, nil)
}

type Users struct {
	CreateUserView *views.View
}
