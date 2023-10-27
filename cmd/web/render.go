package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

type templData struct {
	stringMap  map[string]string
	intMap     map[string]int
	floatMap   map[string]float32
	data       map[string]interface{}
	csrfToken  string
	flash      string
	warning    string
	err        string
	isAuthed   bool
	api        string
	cssVersion string
}

// ErrorResponseMessage represents the structure of the error
// object sent in failed responses.
type ErrorResponseMessage struct {
	Message string `json:"message"`
}

// ErrorResponse represents the structure of the error object sent
// in failed responses.
type ErrorResponse struct {
	Error *ErrorResponseMessage `json:"error"`
}

var funcs = template.FuncMap{}

//go:embed templates
var templateFS embed.FS

func (app *application) addDefaultData(td *templData, r *http.Request) *templData {
	return td
}

func (app *application) renderTmpl(w http.ResponseWriter, r *http.Request, tmplName string, td *templData, partials ...string) error {
	var tmpl *template.Template
	var err error
	tmplPath := fmt.Sprintf("templates/%s.page.tmpl", tmplName)

	// don't read template from cache in development mode.
	if tc, ok := app.tmplCache[tmplPath]; ok && app.config.env == "production" {
		tmpl = tc
	} else {
		tmpl, err = app.buildTmpl(tmplName, tmplPath, partials)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templData{}
	}

	td = app.addDefaultData(td, r)

	err = tmpl.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) buildTmpl(tmplName string, tmplPath string, partials []string) (*template.Template, error) {
	var tmpl *template.Template
	var err error

	// map over partials, replacing partial name with full path to partial.
	if len(partials) > 0 {
		for i, pt := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.tmpl", pt)
		}
	}

	if len(partials) > 0 {
		tmpl, err = template.New(fmt.Sprintf("%s.page.tmpl", tmplName)).Funcs(funcs).ParseFS(templateFS, "templates/base.layout.tmpl", strings.Join(partials, ","), tmplPath)
	} else {
		tmpl, err = template.New(fmt.Sprintf("%s.page.tmpl", tmplName)).Funcs(funcs).ParseFS(templateFS, "templates/base.layout.tmpl", tmplPath)
	}

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.tmplCache[tmplPath] = tmpl
	return tmpl, nil
}

func (app *application) writeJSON(w http.ResponseWriter, v interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		app.errorLog.Println(err)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		app.errorLog.Println(err)
		return err
	}
	return nil
}

func (app *application) writeJSONError(w http.ResponseWriter, v interface{}, code int) {
	w.WriteHeader(code)
	app.writeJSON(w, v)
}

func (app *application) writeJSONErrorMessage(w http.ResponseWriter, message string, code int) {
	resp := &ErrorResponse{
		Error: &ErrorResponseMessage{
			Message: message,
		},
	}
	app.writeJSONError(w, resp, code)
}
