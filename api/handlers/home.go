package api

import (
	"net/http"
)

func handleHomeView(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./static"))

	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./static/index.html")
		return
	}

	fs.ServeHTTP(w, r)
}
