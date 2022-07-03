package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "PARAM")
	_, err := w.Write([]byte(fmt.Sprintf("Hello, %v", param)))
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func BadHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(500), 500)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	_, err = w.Write([]byte(fmt.Sprintf("I got message:\n%v", string(body))))
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header

	a, err := strconv.Atoi(h.Get("a"))
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	b, err := strconv.Atoi(h.Get("b"))
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	w.Header().Add("a+b", strconv.Itoa(a+b))
}
