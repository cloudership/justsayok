package main

import (
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", helloWorld)
	var port, ok = os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}
