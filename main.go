package main

import (
	"database/sql"
	"fmt"
	"log"
_	"github.com/lib/pq"
	"net/http"
	"os"
	"time"
)

var db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func waitHandler(w http.ResponseWriter, r *http.Request) {
	waitTime := 2 * time.Second
	log.Println("waiting for", waitTime)
	time.Sleep(waitTime)

	fmt.Fprintf(w, "Thanks for waiting\n")
}

func sqlHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	err := db.QueryRow("SELECT 1").Scan(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, result)
}

func DBSetup(){
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://localhost:5432/postgres?sslmode=disable"
	}
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func main() {
	DBSetup()

	http.HandleFunc("/wait", waitHandler)
	http.HandleFunc("/list", sqlHandler)
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Running on:", port, "...")
	http.ListenAndServe(":"+port, nil)
}
