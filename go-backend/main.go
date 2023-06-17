package main

import (
	"go-backend/config"
	"log"
)

func main() {
	// get env.json
	config.EnvironmentSetup()
	config.Connect()

	log.Default().Print("\tend\n")
}
