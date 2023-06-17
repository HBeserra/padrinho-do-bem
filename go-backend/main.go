package main

import (
	"go-backend/config"
	"log"
)

func main() {
	// get env.json
	config.EnvironmentSetup()
	config.Connect()
	config.DBConnection.Ping()

	log.Default().Print("\tend\n")
}
