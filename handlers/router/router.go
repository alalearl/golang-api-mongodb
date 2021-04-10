package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jay-supakorn/api-service/app/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var Router *gin.Engine

func Get(client *mongo.Database) *gin.Engine {
	Router := gin.Default()
	Router.Use(Cors())

	userHandler := user.Handler{DB: client}
	Router.GET("/v1/users", userHandler.Lists)
	Router.POST("/v1/user/create", userHandler.CreateNew)

	return Router
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
