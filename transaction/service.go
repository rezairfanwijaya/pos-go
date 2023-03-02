package transaction

import (
	"fmt"
	"pos/helper"
	"time"
)

type IService interface {
	CreateTransaction(input InputTransaction) (Transaction, error)
}

type service struct {
	transactionRepo IRepository
}

func NewService(transactionRepo IRepository) *service {
	return &service{transactionRepo}
}

func (s *service) CreateTransaction(input InputTransaction) (Transaction, error) {
	// assign input to model transaction
	var transaction Transaction
	transaction.Amount = input.Amount
	transaction.Notes = input.Notes
	transaction.Type = input.Type

	date, err := helper.TimeParser(time.RFC822, input.Date)
	if err != nil {
		return transaction, err
	}

	transaction.Date = date

	transactionSaved, err := s.transactionRepo.Save(transaction)
	if err != nil {
		return transaction, fmt.Errorf(
			"failed save new transaction: %v",
			err.Error(),
		)
	}

	return transactionSaved, nil
}
