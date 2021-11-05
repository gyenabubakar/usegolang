package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprintf(w, "<h1>Couldn't find the path \"%s\".</h1>", path)
	HandleError(err, "Couldn't process request. :(")
}

func main() {
	var err error

	homeTemplate, err = template.ParseFiles(
		"views/index.gohtml",
		"views/components/footer.gohtml",
	)
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles(
		"views/contact.gohtml",
		"views/components/footer.gohtml",
	)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	_ = http.ListenAndServe(":8080", router)
}
