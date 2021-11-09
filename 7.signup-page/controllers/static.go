package controllers

import "signup_page/views"

func StaticController() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "views/static/index.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}

type Static struct {
	Home *views.View
	Contact *views.View
}
