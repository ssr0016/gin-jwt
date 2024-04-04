package controller

import (
	"gin-jwt/data/response"
	"gin-jwt/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUserController(repository repository.UserRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	users := controller.userRepository.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Successfully fetch all user details!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
