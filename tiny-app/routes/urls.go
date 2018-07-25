package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"../utils"
)

var urls = []URL{URL{"b2xVn2", "http://lighthouselabs.ca"}}

// URL ...
type URL struct {
	ShortURL string
	LongURL  string
}

// UrlsHandler ..
func UrlsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUrl(w, r)
	case "POST":
		createUrl(w, r)
	}
}

// createUrl ...
func createUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := utils.MakeRandomString()
	longUrl := r.FormValue("longUrl")
	urls = append(urls, URL{shortUrl, longUrl})
	fmt.Println(shortUrl, longUrl)
	http.Redirect(w, r, "/urls", 301)
}

// getUrl ...
func getUrl(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../views/urls_index.html")

	if err != nil {
		log.Printf("You fucked shit up")
	}

	Urls := []URL{URL{"b2xVn2", "http://lighthouselabs.ca"}}
	tmpl.Execute(w, Urls)
}

// redirectToUrls ...
func RedirectToUrls(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/urls", 302)
}
