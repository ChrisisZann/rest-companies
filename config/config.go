package config

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"sync"
	"xm-companies/repository"
)

type Application struct {
	Models *repository.Models
	JwtKey []byte
}

var instance *Application
var once sync.Once
var db *sql.DB

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {

	once.Do(func() {
		instance = &Application{
			Models: repository.New(db),
			JwtKey: []byte("default_secret_key"),
		}
	})
	return instance
}

func (a *Application) LoadConfig(configFile string) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	var dat map[string]interface{}
	json.Unmarshal(data, &dat)

	jwtKey, ok := dat["jwt_key"].(string)
	if !ok {
		log.Fatalf("jwt_key is not a string")
	}

	a.JwtKey = []byte(jwtKey)

	// Bad practise :)
	log.Println("loaded jwt key:", a.JwtKey)
}
