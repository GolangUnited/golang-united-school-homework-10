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

// CatchAllHandler /** Handles all requests.
// If request is not handled by any other handler,
// then this handler will be called.
// It will return 200 status code.
func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// BadStatusHandler handles "/bad" endpoint.
// It will return 500 status code.
func BadStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

// NameHandler handles "/name/{PARAM}" endpoint.
// It writes "Hello, {PARAM}!" to response.
// It returns 200 status code.
func NameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["PARAM"]

	fmt.Fprintf(w, "Hello, %s!", param)
}

// HeaderHandler handles "/headers" endpoint.
// It receives two headers: a and b.
// It adds a+b to the response header.
// a+b is the sum of a and b.
func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	aValue, _ := strconv.Atoi(a)
	bValue, _ := strconv.Atoi(b)

	sum := aValue + bValue
	w.Header().Add("a+b", strconv.Itoa(sum))
	// w.Header().Set("a+b", strconv.Itoa(sum))

}

// DataHandler handles "/data" endpoint.
// It receives body of the request and writes it to response.
// It writes "I got message: {BODY}" to response.
func DataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "I got message:\n%s", string(body))
}

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.
main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", NameHandler).Methods("GET")
	router.HandleFunc("/bad", BadStatusHandler).Methods("GET")
	router.HandleFunc("/data", DataHandler).Methods("POST")
	router.HandleFunc("/headers", HeaderHandler).Methods("POST")
	router.HandleFunc("/", CatchAllHandler).Methods("GET")

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
