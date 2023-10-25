package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/vt", app.virtualTerminal)
	r.Get("/public/", app.public)
	// r.Handle("/public", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	return r
}
