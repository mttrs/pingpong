package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func waitHandler(w http.ResponseWriter, r *http.Request) {
	waitTime := 2 * time.Second
	fmt.Println("waiting for", waitTime)
	time.Sleep(waitTime)
	fmt.Fprintf(w, "Thanks for waiting\n")
}

func main() {
	http.HandleFunc("/wait", waitHandler)
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Running on:", port, "...")
	http.ListenAndServe(":"+port, nil)
}
