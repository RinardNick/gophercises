package handler

import (
	"net/http"
)

// MapHandler will handle mapping urls
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r  *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}