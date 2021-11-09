package controllers

import "signup_page/views"

func StaticController() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "static/index"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

type Static struct {
	Home *views.View
	Contact *views.View
}
