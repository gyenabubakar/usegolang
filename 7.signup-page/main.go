package main

import (
	"first-views/controllers"
	"first-views/views"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView.Render(w, nil, nil)
}

func contact(w http.ResponseWriter, _ *http.Request) {
	contactView.Render(w, nil, nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprintf(w, "<h1>Couldn't find the path \"%s\".</h1>", path)
	HandleError(err, "Couldn't process request. :(")
}

func main() {
	homeView = views.NewView("bootstrap", "views/index.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	usersController := controllers.UsersController()

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/signup", usersController.RenderSignupView).
		Methods("GET")
	router.HandleFunc("/signup", usersController.HandlerUserCreation).
		Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	fmt.Println("Server started!")
	_ = http.ListenAndServe(":8080", router)
}
