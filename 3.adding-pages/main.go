package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func faqs(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "<h1>FAQs Page</h1>")
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
	router.HandleFunc("/faqs", faqs)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	_ = http.ListenAndServe(":8080", router)
}
