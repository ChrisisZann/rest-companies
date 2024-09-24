package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type dbConfig struct {
	db_user     string
	db_password string
	db_host     string
	db_name     string
}

func connectToDB(conf dbConfig) *sql.DB {

	options := "sslmode=disable"
	dataSourceName := "postgres://" + conf.db_user + ":" + conf.db_password + "@" + conf.db_host + "/" + conf.db_name + "?" + options
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("error connecting to postgres database", err)
	}
	return db
}
