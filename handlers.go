package main

import "net/http"

// temporary
func index(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "index", "layout", "mobile", "navbar", "pager", "sidebar")
}
