package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is now better")
}

func main() {
	// Assigning our handleFunc to the "/" pattern
	// for the default mux
	http.HandleFunc("/", handleFunc)
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
