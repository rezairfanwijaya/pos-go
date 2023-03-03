package main

import (
	"log"
	"net/http"
	"pos/auth"
	"pos/database"
	"pos/handler"
	"pos/helper"
	"pos/transaction"
	"pos/user"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
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

	// transaction
	transactionRepo := transaction.NewRepository(connection)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := handler.NewHandlerTransaction(transactionService)

	// init service
	r := gin.Default()
	r.Use(cors.Default())

	// api versioning
	apiV1 := r.Group("/api/v1")

	// route
	apiV1.POST("/login", userHandler.Login)
	apiV1.POST("/transaction/create", authMiddleware(*authService, *userService), transactionHandler.NewTransaction)
	apiV1.PUT("/transaction/update/:id", authMiddleware(*authService, *userService), transactionHandler.UpdateTransactionByID)
	apiV1.DELETE("/transaction/delete/:id", authMiddleware(*authService, *userService), transactionHandler.DeleteTransactionByID)
	apiV1.GET("/transaction/show", authMiddleware(*authService, *userService), transactionHandler.GetAllTransactions)

	env, err := helper.GetENV(".env")
	if err != nil {
		log.Fatal(err)
	}

	if err := r.Run(env["DOMAIN"]); err != nil {
		log.Fatal("error start server", err)
	}
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get header authorization
		authHeader := c.GetHeader("Authorization")

		// must contain bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.GenerateResponse(
				http.StatusUnauthorized,
				"Unauthorized",
				"Access denied",
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get jwt
		tokenString := ""
		headerSplit := strings.Split(authHeader, " ")
		if len(headerSplit) == 2 {
			tokenString = headerSplit[1]
		}

		// jwt validate
		token, err := authService.VerifyTokenJWT(tokenString)
		if err != nil {
			response := helper.GenerateResponse(
				http.StatusUnauthorized,
				"Unauthorized",
				err.Error(),
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get payload
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.GenerateResponse(
				http.StatusUnauthorized,
				"Unauthorized",
				"Invalid Token",
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get userid from payload
		userID := int(claim["user_id"].(float64))

		// get user by userid in token
		userByUserID, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.GenerateResponse(
				http.StatusUnauthorized,
				"Unauthorized",
				err.Error(),
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set context
		c.Set("currentUser", userByUserID)
	}
}
