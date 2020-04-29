package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

// Used to create a new Users controller.
// It will panic if the templates aren't
// parsed correctly.
// Should only be used during initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// Used to render the form where a user can
// create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Used to process the signup form when a user
// submits it. This is used to create a new user
// account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Temp response")
}
