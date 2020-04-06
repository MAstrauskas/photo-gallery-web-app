package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeTemplate *template.Template
var contactTemplate *template.Template

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

func faq(w http.ResponseWriter, r *http.Request) {
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
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")

	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles("views/contact.gohtml")

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(notFound)

	_ = http.ListenAndServe(":3000", router)
}

