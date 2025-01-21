package main

import (
	"log"
	"os"
	"testing"
	"xm-companies/config"
)

var testApp api

func TestMain(m *testing.M) {

	log.Println("initializing testApp")

	initTestConfig, err := config.New("test")
	if err != nil {
		log.Fatal("failed to init config")
	}

	testApp = api{
		cfg: initTestConfig,
	}

	os.Exit(m.Run())

}
