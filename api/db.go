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

	// Capture connection properties.
	// cfg := mysql.Config{
	// 	User:   "postgres",
	// 	Passwd: "4tE_pale",
	//  Net:    "tcp",
	// 	Addr:   "192.168.1.5",
	// 	DBName: "chrisis_home",
	// }

	// Replace the connection string with your actual PostgreSQL connection string
	// db, err := sql.Open("postgres", "postgres://postgres:4tE_pale@192.168.1.5/chrisis_home?sslmode=disable")

	options := "sslmode=disable"
	dataSourceName := "postgres://" + conf.db_user + ":" + conf.db_password + "@" + conf.db_host + "/" + conf.db_name + "?" + options
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("error connecting to postgres database", err)
	}
	return db
}
