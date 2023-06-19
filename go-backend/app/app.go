package app

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func CreateApp() *gin.Engine {
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}
	Server := gin.Default()

	Server.Use(CORSMiddleware())
	Server.Use(gzip.Gzip(gzip.DefaultCompression))

	Server.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{
			"error": "resource not found",
		})
	})

	return Server
}

func StartServer(server *gin.Engine) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port <= 0 || port >= 65535 {
		port = 3000
	}

	err := server.Run(fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}
}
