package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/GolangUnited/helloweb/internal/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	r := chi.NewRouter()

	setupMiddleware(r)

	r.Get("/name/{PARAM}", controllers.NameHandler)
	r.Get("/bad", controllers.BadHandler)
	r.Post("/data", controllers.DataHandler)
	r.Post("/headers", controllers.HeadersHandler)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r); err != nil {
		panic(err)
	}
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
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
