package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Detected %d CPU cores.\n", numCPU)

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Running on:", port, "...")

	http.ListenAndServe(":"+port, nil)
}
