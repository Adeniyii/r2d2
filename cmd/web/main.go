package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
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

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Printf("Starting HTTP server on port %d in %s mode\n", app.config.port, app.config.env)

	return srv.ListenAndServe()
}

const (
	STRIPE_KEY    = "STRIPE_KEY"
	STRIPE_SECRET = "STRIPE_SECRET"
)

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Port to host frontend on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment <dev|prod>")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to API")

	flag.Parse()

	if sk, ok := os.LookupEnv(STRIPE_KEY); !ok {
		log.Fatalf("missing env: %s\n", STRIPE_KEY)
	} else {
		cfg.stripe.key = sk
	}

	if ss, ok := os.LookupEnv(STRIPE_SECRET); !ok {
		log.Fatalf("missing env: %s\n", STRIPE_SECRET)
	} else {
		cfg.stripe.secret = ss
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	tc := make(map[string]*template.Template)

	application := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
	}
}
