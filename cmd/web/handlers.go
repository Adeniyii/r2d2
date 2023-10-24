package main

import "net/http"

func (app *application) virtualTerminal(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Hit virtual terminal")
}
