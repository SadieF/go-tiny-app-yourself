package routes

import (
	"html/template"
	"log"
	"net/http"
)

// URL ...
type URL struct {
	ShortURL string
	LongURL  string
}

// GetUrls ...
func GetUrls(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../views/urls_index.html")

	if err != nil {
		log.Printf("You fucked shit up")
	}

	Urls := []Url{Url{"b2xVn2", "http://lighthouselabs.ca"}}
	tmpl.Execute(w, Urls)
}
