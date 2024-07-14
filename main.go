package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/qinains/fastergoding"
	"github.com/zi-bot/simple-gin-rest/config"
	"github.com/zi-bot/simple-gin-rest/middleware"
	"github.com/zi-bot/simple-gin-rest/routes"
	"github.com/zi-bot/simple-gin-rest/utils/logger"
)

func main() {
	r := gin.Default()
	if os.Getenv("GIN_MODE") != "release" {
		fastergoding.Run()
	}

	db := config.ConnectDatabase()

	logger := logger.GetLogger()
	logger.Info("Starting the application...")

	r.Use(
		middleware.Recover(),
		middleware.Cors(),
		middleware.RateLimitter(),
	)

	routes.SetupRoutes(r, db)

	r.Run(":8080")
}
