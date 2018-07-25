package main

import (
	"math/rand"
	"net/http"
	"time"

	"./routes"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", routes.RedirectToUrls)
	http.HandleFunc("/urls", routes.UrlsHandler)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

}
