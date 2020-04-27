package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	must(contactView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	must(signupView.Render(w, nil))
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
	signupView = views.NewView("bootstrap", "views/signup.gohtml")

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/signup", signup)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(notFound)

	_ = http.ListenAndServe(":3000", router)
}

// For development purposes only
func must(err error) {
	if err != nil {
		panic(err)
	}
}
