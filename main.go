package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set the HTTP header
	// Set the Content-Type HTTP header
	w.Header().Set("Content-Type", "text/html; charset=utf=8")
	fmt.Fprintf(w, "this is now better")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	// Set the HTTP header
	w.Header().Set("Content-Type", "text/html; charset=utf=8")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>Reach out to me on <a href=\"https://linkedin.com/in/planetrobbie\">LinkedIn</a>.")
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	// Set the HTTP header status code 404
	w.WriteHeader(404)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// handle page not found error
		notfoundHandler(w, r)
	}
}

func main() {
	// Assigning our handleFunc to the "/" pattern
	// for the default mux
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Server starting on Port :3000")

	// listen on 127.0.0.1 to avoid macos popup asking for auth
	// otherwise sign your binary: codesign -s - <binary>
	err := http.ListenAndServe("127.0.0.1:3000", nil)

	// checking errors isn't stricly necessary at main end
	// added just for learning purpose
	if err != nil {
		panic("can't listen")
	}
}
