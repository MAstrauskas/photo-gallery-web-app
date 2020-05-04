package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// interface{} - can be any type
func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}
