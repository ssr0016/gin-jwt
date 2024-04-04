package router

import (
	"gin-jwt/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(authenticationController *controller.AuthenticationController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})

	router := service.Group("/api")

	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	return service
}
