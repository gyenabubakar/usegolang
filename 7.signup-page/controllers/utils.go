package controllers

import (
	"github.com/gorilla/schema"
	"net/http"
)

func ParseForm(r *http.Request, d interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	decoder := schema.NewDecoder()
	if err := decoder.Decode(d, r.PostForm); err != nil {
		return err
	}
	return nil
}
