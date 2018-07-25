package main

import (
	"net/http"

	"./routes"
)

func main() {
	http.HandleFunc("/urls", routes.GetUrls)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

}
