package main

import (
	"fmt"
	"log"
	"net/http"
	"signup_page/controllers"
	"signup_page/views"

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
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/index.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	usersController := controllers.UsersController()

	router := mux.NewRouter()

	// serve static files
	router.PathPrefix("/static/").
		Handler(staticFilesHandler()).
		Methods("GET")

	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact).
		Methods("GET")

	// /signup controllers
	router.HandleFunc("/signup", usersController.RenderSignupView).
		Methods("GET")
	router.HandleFunc("/signup", usersController.HandlerUserCreation).
		Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func staticFilesHandler() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir("static/")))
}
