package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
)

var db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func waitHandler(w http.ResponseWriter, r *http.Request) {
	t := os.Getenv("WAIT_TIME")
	if t == "" {
		t = "2s"
	}
	waitTime, _ := time.ParseDuration(t)
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

func primeHander(w http.ResponseWriter, r *http.Request) {
	max := 1000000
	primes := make([]int64, 0)
	for {
		for n := 2; n <= max; n++ {
			flag := true
			for m := 2; m < n; m++ {
				if (n % m) == 0 {
					flag = false
					break
				}
			}
			if flag {
				primes = append(primes, int64(n))
			}
		}
	}
}

func DBSetup() {
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

func pingHandler(w http.ResponseWriter, r *http.Request) {
	c, err := redis.DialURL(os.Getenv("REDIS_URL"))
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	for i := 0; i < 10; i++ {
		s, err := c.Do("PING")
		if err != nil {
			fmt.Println(err)
		}
		if i == 0 {
			fmt.Println(s)
		}
	}
	fmt.Fprintf(w, "PONG")
}

func acmHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "mtT6rvZnH5bNa8BmrIiZFue-gSUJf71IbTPaY6ikBSk.dJa0PtqeEKVpkuRerpQSHtPj7iCJKLFZlVsrmIm6res")
}

func main() {
	// TODO: https://gist.github.com/tsenart/5fc18c659814c078378d
	//	DBSetup()

	http.HandleFunc("/.well-known/acme-challenge/mtT6rvZnH5bNa8BmrIiZFue-gSUJf71IbTPaY6ikBSk", acmHandler)
	http.HandleFunc("/wait", waitHandler)
	//	http.HandleFunc("/list", sqlHandler)
	http.HandleFunc("/prime", primeHander)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Running on:", port, "...")
	http.ListenAndServe(":"+port, nil)
}
