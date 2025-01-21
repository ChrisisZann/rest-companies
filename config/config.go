package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"xm-companies/repository"

	"github.com/go-sql-driver/mysql"
)

type Application struct {
	Models *repository.Models
	JwtKey []byte
	SqlCfg *mysql.Config
	Logger *log.Logger
}

type inputConfig struct {
	JwtKey []byte
	SqlCfg *mysql.Config
}

var instance *Application
var once sync.Once
var db *sql.DB

func New(configFile string) (*Application, error) {
	inputConfig := LoadConfig(configFile)
	return GetInstance(inputConfig)
}

func GetInstance(ic inputConfig) (*Application, error) {

	err := os.MkdirAll("/var/log/xm-companies/", 0755)
	if err != nil {
		fmt.Println("Failed to create log directory")
		return nil, err
	}

	logFile, err := os.Create("/var/log/xm-companies/service.log")
	if err != nil {
		fmt.Println("Failed to create log file")
		return nil, err
	}

	once.Do(func() {
		instance = &Application{
			Models: repository.New(connectToDB(*ic.SqlCfg)),
			JwtKey: ic.JwtKey,
			Logger: log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile),
		}
	})

	return instance, nil
}

func LoadConfig(configFile string) inputConfig {
	var loadConfig inputConfig

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
	loadConfig.JwtKey = []byte(jwtKey)

	db_user, ok := dat["db_user"].(string)
	if !ok {
		log.Fatalf("db_user is not a string")
	}

	db_password, ok := dat["db_password"].(string)
	if !ok {
		log.Fatalf("db_password is not a string")
	}

	db_host, ok := dat["db_host"].(string)
	if !ok {
		log.Fatalf("db_host is not a string")
	}

	db_name, ok := dat["db_name"].(string)
	if !ok {
		log.Fatalf("db_name is not a string")
	}

	log.Println("db user: " + db_user)
	log.Println("db pass: " + db_password)
	log.Println("db host: " + db_host)
	log.Println("db name: " + db_name)

	loadConfig.SqlCfg = &mysql.Config{
		User:   db_user,
		Passwd: db_password,
		Net:    "tcp",
		Addr:   db_host,
		DBName: db_name,
	}
	// Debug only, Bad practise :)
	// log.Println("loaded jwt key:", a.JwtKey)

	return loadConfig
}
