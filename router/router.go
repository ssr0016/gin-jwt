package router

import (
	"gin-jwt/controller"
	"gin-jwt/middleware"
	"gin-jwt/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userREpository repository.UserRepository, authenticationController *controller.AuthenticationController, userController *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})

	router := service.Group("/api")

	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	userRouter := router.Group("/users")
	userRouter.GET("", middleware.DeserializeUser(userREpository), userController.GetUsers)

	return service
}
