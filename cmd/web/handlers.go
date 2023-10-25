package main

import "net/http"

func (app *application) virtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTmpl(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) public(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("hit css route")
	app.infoLog.Println(r.Header)
	http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
}
