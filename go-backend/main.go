package main

import (
	"github.com/gin-gonic/gin"
	"go-backend/app"
	"go-backend/config"
	"log"
	"net/http"
)

func main() {
	// Obtem a configuração e as conexões com os recursos
	config.EnvironmentSetup()
	config.Connect()
	config.DBConnection.Ping()

	// Inicia o servidor web
	server := app.CreateApp()
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "online",
		})
	})

	app.StartServer(server)

	log.Default().Print("\tend\n")
}
