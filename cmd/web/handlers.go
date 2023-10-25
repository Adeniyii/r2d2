package main

import "net/http"

func (app *application) virtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTmpl(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}
