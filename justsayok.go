package main

import (
	"net/http"
	"os"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {})
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	fmt.Printf(`{"event": "started", "message": "Server started on port %s"}%s`, port, "\n")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(fmt.Errorf(`{"event": "error_starting", "message": "%v"}`, err))
		os.Exit(1)
	}
}
