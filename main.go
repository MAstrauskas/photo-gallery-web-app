package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	_, _ = fmt.Fprint(w, "<h1>Welcome to the site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	_, _ = fmt.Fprint(w, "To get in touch, please send an email to "+
		"<a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	_ = http.ListenAndServe(":3000", r)
}

