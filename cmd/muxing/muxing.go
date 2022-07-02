package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.

func handleNameGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	rv := mux.Vars(r)
	_, err := fmt.Fprintf(w, "Hello, %s!", rv["PARAM"])
	if err != nil {
		log.Fatal(err)
	}
}

func handleBadGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func handleDataPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "I got message:\n%s", string(body))
	if err != nil {
		log.Fatal(err)
	}
}

func handleHeaderPost(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		a = 0
		fmt.Println("Incorrect input")
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		b = 0
		fmt.Println("Incorrect input")
	}
	w.Header().Add("a+b", strconv.Itoa(a+b))
}

func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", handleNameGet).Methods("GET")
	router.HandleFunc("/bad", handleBadGet).Methods("GET")
	router.HandleFunc("/data", handleDataPost).Methods("POST")
	router.HandleFunc("/headers", handleHeaderPost).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}

	Start(host, port)
}
