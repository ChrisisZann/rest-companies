package config

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func connectToDB(conf mysql.Config) *sql.DB {

	options := "sslmode=disable"
	dataSourceName := "postgres://" + conf.User + ":" + conf.Passwd + "@" + conf.Addr + "/" + conf.DBName + "?" + options
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("error connecting to postgres database", err)
	}
	return db
}
