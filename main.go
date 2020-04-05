package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	_, _ = fmt.Fprint(w, "<h1>Welcome to the site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	_, _ = fmt.Fprint(w, "To get in touch, please send an email to "+
		"<a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

	_, _ = fmt.Fprint(w, "<h1>Frequently Asked Questions</h1>" +
		"<h2>What is this website about?</h2>" +
		"<p>It's about learning to build real world applications using Go (Golang)</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	_, _ = fmt.Fprint(w, "<h1>404 page not found<h1>" +
		"<h2>The page you were trying to access does not exist.</h2>")
}

func main() {
	router := httprouter.New()

	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/faq", faq)
	router.NotFound = http.HandlerFunc(notFound)

	_ = http.ListenAndServe(":3000", router)
}

