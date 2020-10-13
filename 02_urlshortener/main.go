package main

import (
	"fmt"
	"handler"
	"net/http"

	"github.com/RinardNick/gophercises/02_urlshortener/handler"
)

func main() {
	mux := defaultMux()

	// Build MapHandler using mux as fallback
	pathsToUrls := map[string]string{
		"/short-beta":    "https://beta.talagentfinancial.com",
		"/short-staging": "https://staging.talagentfinancial.com",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
