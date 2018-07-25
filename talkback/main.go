package main

import (
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

}

// Say Hello ...
func SayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}
