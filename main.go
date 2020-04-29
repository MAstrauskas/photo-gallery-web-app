package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/controllers"

	"lenslocked.com/views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	must(contactView.Render(w, nil))
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

	usersC := controllers.NewUsers()

	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/contact", contact).Methods("GET")
	router.HandleFunc("/faq", faq).Methods("GET")
	router.HandleFunc("/signup", usersC.New).Methods("GET")
	router.HandleFunc("/signup", usersC.Create).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(notFound)

	_ = http.ListenAndServe(":3000", router)
}

// For development purposes only
func must(err error) {
	if err != nil {
		panic(err)
	}
}
