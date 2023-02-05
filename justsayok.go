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
	println(`{"event": "started", "message": "Server started on port ` + port + `"}`)
	http.ListenAndServe(":"+port, nil)
}
