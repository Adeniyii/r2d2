package main

import (
	"net/http"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func (app *application) virtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTmpl(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) stripeConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	err := app.writeJSON(w, struct {
		PublishableKey string `json:"publishableKey"`
	}{
		PublishableKey: app.config.stripe.key,
	})
	if err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) createPaymentIntent(w http.ResponseWriter, r *http.Request) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok {
			app.errorLog.Printf("Stripe error: %v\n", stripeErr.Error())
			app.writeJSONErrorMessage(w, stripeErr.Error(), http.StatusBadRequest)
		} else {
			app.errorLog.Printf("Other error: %v\n", err.Error())
			app.writeJSONErrorMessage(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	app.writeJSON(w, struct {
		ClientSecret string `json:"clientSecret"`
	}{
		ClientSecret: pi.ClientSecret,
	})
}
