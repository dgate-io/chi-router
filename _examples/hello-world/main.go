package main

import (
	"net/http"

	"github.com/dgate-io/chi-router"
	"github.com/dgate-io/chi-router/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":3333", r)
}
