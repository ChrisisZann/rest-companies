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

	testApp = api{
		cfg: config.New(nil),
	}

	os.Exit(m.Run())

}
