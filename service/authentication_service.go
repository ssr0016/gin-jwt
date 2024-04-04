package service

import (
	"gin-jwt/data/request"
)

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest)
}
