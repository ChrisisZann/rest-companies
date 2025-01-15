package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"xm-companies/config"
	"xm-companies/events"
)

const port = ":8888"

var cfgFile = flag.String("c", "config.json", "configuration file")

type api struct {
	cfg               *config.Application
	hub               *events.Hub
	internalPublisher *events.InternalPublisher
}

func main() {

	flag.Parse()
	log.Println("Input config:", *cfgFile)

	companies_api := api{
		cfg: config.New(*cfgFile),
		hub: events.NewHub(),
	}

	companies_api.internalPublisher = events.NewPublisher(companies_api.hub)
	go companies_api.hub.Run()

	srv := &http.Server{
		Addr:              port,
		Handler:           companies_api.router(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	log.Println("Starting web application on port", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
