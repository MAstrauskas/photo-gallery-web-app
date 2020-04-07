package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	_, _ = fmt.Fprint(w, "<h1>Frequently Asked Questions</h1>"+
		"<h2>What is this website about?</h2>"+
		"<p>It's about learning to build real world applications using Go (Golang)</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	_, _ = fmt.Fprint(w, "<h1>404 page not found<h1>"+
		"<h2>The page you were trying to access does not exist.</h2>")
}

func main() {
	homeView = views.NewView(
		"bootstrap",
		"views/home.gohtml",
	)
	contactView = views.NewView(
		"bootstrap",
		"views/contact.gohtml",
	)

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(notFound)

	_ = http.ListenAndServe(":3000", router)
}
