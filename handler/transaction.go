package handler

import (
	"fmt"
	"net/http"
	"pos/helper"
	"pos/transaction"
	"strconv"

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
	var input transaction.InputNewTransaction

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

func (h *transactionHandler) GetAllTransactions(c *gin.Context) {
	amount := c.DefaultQuery("amount", "amount desc")
	date := c.DefaultQuery("date", "date desc")
	transactionType := c.DefaultQuery("type", "e")
	fromAmount := c.DefaultQuery("from", "1000")
	toAmount := c.DefaultQuery("to", "9999999999999999")
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")

	fromAmountNumber, err := helper.ConvertStringToInt(fromAmount)
	if err != nil {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"failed to convert string",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	toAmountNumber, err := helper.ConvertStringToInt(toAmount)
	if err != nil {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"failed to convert string",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	limitNumber, err := helper.ConvertStringToInt(limit)
	if err != nil {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"failed to convert string",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	pageNumber, err := helper.ConvertStringToInt(page)
	if err != nil {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"failed to convert string",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	params := transaction.ParamsGetAllTransaction{
		Amount:          amount,
		Date:            date,
		TransactionType: transactionType,
		FromAmount:      fromAmountNumber,
		ToAmount:        toAmountNumber,
		Limit:           limitNumber,
		Page:            pageNumber,
	}

	path := c.Request.URL.Path
	url := fmt.Sprintf("http://localhost:2222%v", path)

	transactionsWithPagination, err := h.transactionService.GetAllTransaction(params, url)
	if err != nil {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"failed to get transactions",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.GenerateResponse(
		http.StatusOK,
		"success",
		transactionsWithPagination,
	)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) UpdateTransactionByID(c *gin.Context) {
	id := c.Param("id")

	transactionID, err := strconv.Atoi(id)
	if err != nil || transactionID < 1 {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"invalid id transaction",
			"id must be int and grather than 0",
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// binding input
	var input transaction.InputEditTransaction
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

	transactionUpdated, err := h.transactionService.UpdateTransaction(input, transactionID)
	if err != nil {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"failed to update transaction",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.GenerateResponse(
		http.StatusOK,
		"success",
		transactionUpdated,
	)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) DeleteTransactionByID(c *gin.Context) {
	id := c.Param("id")

	transactionID, err := strconv.Atoi(id)
	if err != nil || transactionID < 1 {
		response := helper.GenerateResponse(
			http.StatusBadRequest,
			"invalid id transaction",
			"id must be int and grather than 0",
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := h.transactionService.DeleteTransactionByID(transactionID); err != nil {
		response := helper.GenerateResponse(
			http.StatusNotFound,
			"failed delete transaction",
			err.Error(),
		)

		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.GenerateResponse(
		http.StatusOK,
		"success",
		"success delete transaction",
	)

	c.JSON(http.StatusOK, response)
}
