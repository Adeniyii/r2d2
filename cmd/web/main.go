package main

import (
	"html/template"
	"log"
)

const version = "1.0.0"
const cssVersion = "1" // Incrementing will force browers invalidate their cache of static files

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}
