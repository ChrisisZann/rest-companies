package config

import (
	"database/sql"
	"encoding/json"
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
	JwtKey  []byte
	SqlCfg  *mysql.Config
	LogFile *os.File
}

var instance *Application
var once sync.Once
var db *sql.DB

func New(configFile string) *Application {
	inputConfig := LoadConfig(configFile)
	return GetInstance(inputConfig)
}

func GetInstance(ic inputConfig) *Application {

	once.Do(func() {
		instance = &Application{
			Models: repository.New(connectToDB(*ic.SqlCfg)),
			JwtKey: ic.JwtKey,
			Logger: log.New(ic.LogFile, "", log.Ldate|log.Ltime|log.Lshortfile),
		}
	})

	return instance
}

func LoadConfig(configFile string) inputConfig {
	var loadConfig inputConfig

	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	var dat map[string]interface{}
	json.Unmarshal(data, &dat)

	logName := "service"
	logFileString, ok := dat["log_dir"].(string)
	if !ok {
		log.Fatal("log_file is not a string")
	}
	err = os.MkdirAll(logFileString, 0755)
	if err != nil {
		log.Fatal("Failed to create log directory")
	}
	logFile, err := os.Create(logFileString + "/" + logName + ".log")
	if err != nil {
		log.Fatal("Failed to create log file")
	}
	loadConfig.LogFile = logFile
	log.Println("log: " + logFileString + "/" + logName + ".log")

	jwtKey, ok := dat["jwt_key"].(string)
	if !ok {
		log.Fatal("jwt_key is not a string")
	}
	loadConfig.JwtKey = []byte(jwtKey)

	db_user, ok := dat["db_user"].(string)
	if !ok {
		log.Fatal("db_user is not a string")
	}

	db_password, ok := dat["db_password"].(string)
	if !ok {
		log.Fatal("db_password is not a string")
	}

	db_host, ok := dat["db_host"].(string)
	if !ok {
		log.Fatal("db_host is not a string")
	}

	db_name, ok := dat["db_name"].(string)
	if !ok {
		log.Fatal("db_name is not a string")
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
