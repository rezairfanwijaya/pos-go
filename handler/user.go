package handler

import (
	"net/http"
	"pos/auth"
	"pos/helper"
	"pos/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.IService
	authService auth.IAuthService
}

func NewHandlerUser(
	userService user.IService,
	authService auth.IAuthService,
) *userHandler {
	return &userHandler{
		userService,
		authService,
	}
}

func (h *userHandler) Login(c *gin.Context) {
	var userInput user.InputUserLogin

	if err := c.BindJSON(&userInput); err != nil {
		errBinding := helper.FormatingErrorBinding(err)
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"invalid input",
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userLoggedin, err := h.userService.Login(userInput)
	if err != nil {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"failed to login",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateTokenJWT(userLoggedin.ID)
	if err != nil {
		response := helper.GenerateResponse(
			http.StatusInternalServerError,
			"failed to generate token",
			err.Error(),
		)

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	userFormatted := user.UserFormatter(userLoggedin, token)

	response := helper.GenerateResponse(
		http.StatusOK,
		"success login",
		userFormatted,
	)

	c.JSON(http.StatusOK, response)
}
