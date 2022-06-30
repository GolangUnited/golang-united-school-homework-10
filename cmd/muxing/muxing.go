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

func NameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Hello, %v!", vars["PARAM"])
	if err != nil {
		_ = fmt.Errorf("[NameHandler] Error: %w", err)
	}
}

func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		_ = fmt.Errorf("[DataHandler] Error: %w", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = io.WriteString(w, fmt.Sprintf("I got message:\n%v", string(b)))
	if err != nil {
		_ = fmt.Errorf("[DataHandler] Error: %w", err)
	}
}

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	a, err := strconv.Atoi(h.Get("a"))
	if err != nil {
		_ = fmt.Errorf("[DataHandler] Error: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := strconv.Atoi(h.Get("b"))
	if err != nil {
		_ = fmt.Errorf("[DataHandler] Error: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sum := a + b

	fmt.Println(fmt.Sprintf("%d", sum))

	w.Header().Add("a+b", fmt.Sprintf("%d", sum))
	w.WriteHeader(http.StatusOK)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	r := mux.NewRouter()

	r.HandleFunc("/name/{PARAM}", NameHandler).Methods("GET")
	r.HandleFunc("/bad", BadHandler).Methods("GET")
	r.HandleFunc("/data", DataHandler).Methods("POST")
	r.HandleFunc("/headers", HeadersHandler).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r); err != nil {
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
