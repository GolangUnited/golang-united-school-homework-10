package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func handleNameParam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params["PARAM"]
	fmt.Fprintf(w, "Hello, "+param+"!")

	return
}

func handleBadParam(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	return
}

func handdleBodyParam(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)

	if err != nil {
		errors.New("something went wrong with http response")
		return
	}

	fmt.Fprintf(w, "I got message:\n"+string(data))

	return
}

func handleHeadersParam(w http.ResponseWriter, r *http.Request) {
	hdrA := r.Header.Get("a")
	hdrB := r.Header.Get("b")

	if hdrA == "" || hdrB == "" {
		errors.New("No available header!")
		return
	}

	res1, e1 := strconv.Atoi(hdrA)
	res2, e2 := strconv.Atoi(hdrB)

	if e1 == nil && e2 == nil {
		w.Header().Set("a+b", strconv.Itoa(res1+res2))
	}

	return
}

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.
main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", handleNameParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBadParam).Methods(http.MethodGet)
	router.HandleFunc("/data", handdleBodyParam).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeadersParam).Methods(http.MethodPost)

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
