package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		_, _ = fmt.Fprint(w, "<h1>Welcome to the site!</h1>")
	} else if r.URL.Path == "/contact" {
		_, _ = fmt.Fprint(w, "To get in touch, please send an email to "+
			"<a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, "<h1>Sorry, the page you requested does not exist.</h1><p>Please email us if you keep being sent to an invalid page.</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	_ = http.ListenAndServe(":3000", nil)
}

