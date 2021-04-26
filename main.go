package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jay-supakorn/api-service/database"
	"github.com/jay-supakorn/api-service/functions/logger"
	"github.com/jay-supakorn/api-service/handlers"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.DebugMode)
	if err := godotenv.Load(); err != nil {
		logger.Info.Println("failed to load env vars")
	}
	client := database.ConnectDB()
	engine := handlers.API(client)
	logger.Info.Println(engine.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT"))))
}
