package main

import (
	"log"
	"pos/auth"
	"pos/database"
	"pos/handler"
	"pos/helper"
	"pos/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	connection, err := database.NewConnection(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	// auth
	authService := auth.NewAuthService()

	// user
	userRepo := user.NewRepository(connection)
	userService := user.NewService(userRepo)
	userHandler := handler.NewHandlerUser(userService, authService)

	// init service
	r := gin.Default()
	r.Use(cors.Default())

	// api versioning
	apiV1 := r.Group("/api/v1")

	// route
	apiV1.POST("/login", userHandler.Login)

	
	env, err := helper.GetENV(".env")
	if err != nil {
		log.Fatal(err)
	}

	if err := r.Run(env["DOMAIN"]); err != nil {
		log.Fatal(err)
	}
}
