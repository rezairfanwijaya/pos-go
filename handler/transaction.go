package handler

import (
	"net/http"
	"pos/helper"
	"pos/transaction"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.IService
}

func NewHandlerTransaction(
	transactionService transaction.IService,
) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) NewTransaction(c *gin.Context) {
	var input transaction.InputTransaction

	if err := c.BindJSON(&input); err != nil {
		errBinding := helper.FormatingErrorBinding(err)
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"invalid input",
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	newTransaction, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"failed to create transaction",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.GenerateResponse(
		http.StatusOK,
		"success",
		newTransaction,
	)

	c.JSON(http.StatusOK, response)

}
