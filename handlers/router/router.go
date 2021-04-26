package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jay-supakorn/api-service/src/controllers/login"
	"github.com/jay-supakorn/api-service/src/controllers/user"
	"github.com/jay-supakorn/api-service/src/services"
	"go.mongodb.org/mongo-driver/mongo"
)

var Router *gin.Engine

func Get(client *mongo.Database) *gin.Engine {
	Router := gin.Default()
	Router.Use(Cors())

	var loginService services.LoginService = services.StaticLoginService()
	var jwtService services.JWTService = services.JWTAuthService()
	var loginController login.LoginController = login.LoginHandler(loginService, jwtService)

	userHandler := user.Handler{DB: client}
	Router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"access_token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

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
