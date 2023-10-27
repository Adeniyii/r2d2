package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/vt", app.virtualTerminal)
	r.Get("/stripe-config", app.stripeConfig)
	r.Get("/create-payment-intent", app.createPaymentIntent)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	return r
}
