package main

import (
	"gin-jwt/config"
	"gin-jwt/controller"
	"gin-jwt/helper"
	"gin-jwt/model"
	"gin-jwt/repository"
	"gin-jwt/router"
	"gin-jwt/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variables", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	// Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)

	// Init Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)

	// Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)

	router := router.NewRouter(authenticationController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
