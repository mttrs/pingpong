package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
	// io.WriteString(w, "pong")
}

func main() {
	http.HandleFunc("/ping", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Running on", port, "...")
	http.ListenAndServe(":"+port, nil)
}
