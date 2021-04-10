package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jay-supakorn/api-service/handlers/router"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Engine *gin.Engine
}

func API(db *mongo.Database) *gin.Engine {
	r := router.Get(db)
	return r
}

func Pong() {
}
