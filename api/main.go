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
	config *config.Application
	hub    *events.Hub
}

func main() {

	flag.Parse()
	log.Println("Input config:", *cfgFile)

	db_conf := dbConfig{
		db_user:     "postgres",
		db_password: "4tE_pale",
		db_host:     "192.168.1.5",
		db_name:     "chrisis_home",
	}

	chr_api := api{
		config: config.New(connectToDB(db_conf)),
		hub:    events.NewHub(),
	}

	chr_api.config.LoadConfig(*cfgFile)

	go chr_api.hub.Run()

	log.Printf("Loaded from cfg file, key: %s", chr_api.config.JwtKey)

	srv := &http.Server{
		Addr:              port,
		Handler:           chr_api.router(),
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
