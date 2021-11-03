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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)

	_ = http.ListenAndServe(":8080", router)
}

//func handleError(err error, message string) {
//	if err != nil {
//		log.Fatal(message)
//	}
//}
