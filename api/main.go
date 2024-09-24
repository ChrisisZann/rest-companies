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
	config            *config.Application
	hub               *events.Hub
	internalPublisher *events.InternalPublisher
}

func main() {

	flag.Parse()
	log.Println("Input config:", *cfgFile)

	// db_conf := dbConfig{
	// 	db_user:     "postgres",
	// 	db_password: "4tE_pale",
	// 	db_host:     "192.168.1.5",
	// 	db_name:     "chrisis_home",
	// }

	// docker postgres
	db_conf := dbConfig{
		db_user:     "postgres",
		db_password: "myverysecretpassword",
		db_host:     "host.docker.internal",
		db_name:     "docker_home",
	}

	companies_api := api{
		config: config.New(connectToDB(db_conf)),
		hub:    events.NewHub(),
	}
	companies_api.internalPublisher = events.NewPublisher(companies_api.hub)
	companies_api.config.LoadConfig(*cfgFile)
	go companies_api.hub.Run()

	log.Println("Loaded jwt key from cfg file")

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
