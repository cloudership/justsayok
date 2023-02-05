package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {})
	var port, ok = os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}
