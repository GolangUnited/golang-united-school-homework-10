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
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{param}", GetParam).Methods("GET")
	router.HandleFunc("/bad", GetErr).Methods("GET")
	router.HandleFunc("/data", PostParam).Methods("POST")
	router.HandleFunc("/headers", PostHeaders).Methods("POST")
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
func GetParam(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	resp.WriteHeader(http.StatusOK)
	fmt.Fprintf(resp, "Hello, %v!", vars["param"])
}

func GetErr(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusInternalServerError)
}

func PostParam(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	body, _ := io.ReadAll(req.Body)
	fmt.Fprintf(resp, "I got message:\n%v", string(body))
}

func PostHeaders(resp http.ResponseWriter, req *http.Request) {
	a, b := req.Header.Get("a"), req.Header.Get("b")
	var intA, intB int
	intA, _ = strconv.Atoi(a)
	intB, _ = strconv.Atoi(b)
	r := intA + intB
	resp.Header().Set("a+b", strconv.Itoa(r))
	resp.WriteHeader(http.StatusOK)
}
