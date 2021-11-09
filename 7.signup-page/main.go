package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"signup_page/controllers"
)

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
	usersController := controllers.UsersController()
	staticController := controllers.StaticController()

	router := mux.NewRouter()

	// serve static files
	router.PathPrefix("/static/").
		Handler(staticFilesHandler()).
		Methods("GET")

	router.Handle("/", staticController.Home).
		Methods("GET")
	router.Handle("/contact", staticController.Contact).
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
