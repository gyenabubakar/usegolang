package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprintf(w, "Welcome to my awesome site!")
	//handleError(err, "Failed to handle request")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprintf(w, "<h1>Couldn't find the path \"%s\".</h1>", path)
	HandleError(err, "Couldn't process request. :(")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	_ = http.ListenAndServe(":8080", router)
}
