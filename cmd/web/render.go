package main

import (
	"embed"
	"fmt"
	"html/template"
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
		tmpl, err = template.New(fmt.Sprintf(tmplPath)).Funcs(funcs).ParseFS(templateFS, "templates/base.layout.tmpl", strings.Join(partials, ","), tmplPath)
	} else {
		tmpl, err = template.New(fmt.Sprintf(tmplPath)).Funcs(funcs).ParseFS(templateFS, "templates/base.layout.tmpl", tmplPath)
	}

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.tmplCache[tmplPath] = tmpl
	return tmpl, nil
}
