package service

import (
	"errors"
	"gin-jwt/config"
	"gin-jwt/data/request"
	"gin-jwt/helper"
	"gin-jwt/model"
	"gin-jwt/repository"
	"gin-jwt/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewAuthenticationServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Login implements AuthenticationService.
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {

	// Find username in database
	new_user, user_err := a.UserRepository.FindByUsername(users.Username)
	if user_err != nil {
		return "", errors.New("invalid username or password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_user.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate Token
	token, err := utils.GenerateToken(config.TokenExpiresIn, new_user, config.TokenSecret)
	helper.ErrorPanic(err)
	return token, nil
}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}

	a.UserRepository.Save(newUser)
}
