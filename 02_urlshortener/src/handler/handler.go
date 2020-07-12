package handler

import (
	"fmt"
	"net/http"
)

// MapHandler will handle mapping urls
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r  *http.Request) {
		path := r.URL.Path
		fmt.Println(path)
		fmt.Println(pathsToUrls[path])
		if dest, ok := pathsToUrls[path]; ok {
			fmt.Println(dest, ok)
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}