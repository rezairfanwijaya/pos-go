package transaction

import (
	"fmt"
	"pos/helper"
)

type IService interface {
	CreateTransaction(input InputNewTransaction) (Transaction, error)
	GetAllTransaction(params ParamsGetAllTransaction, url string) (PaginationTransaction, error)
	UpdateTransaction(input InputEditTransaction, transactionID int) (Transaction, error)
	DeleteTransactionByID(transactionID int) error
}

type service struct {
	transactionRepo IRepository
}

func NewService(transactionRepo IRepository) *service {
	return &service{transactionRepo}
}

func (s *service) CreateTransaction(input InputNewTransaction) (Transaction, error) {
	// assign input to model transaction
	var transaction Transaction
	transaction.Amount = input.Amount
	transaction.Notes = input.Notes
	transaction.Type = input.Type

	date, err := helper.TimeParser(input.Date)
	if err != nil {
		return transaction, err
	}

	transaction.Date = date

	// save
	transactionSaved, err := s.transactionRepo.Save(transaction)
	if err != nil {
		return transaction, fmt.Errorf(
			"failed save new transaction: %v",
			err.Error(),
		)
	}

	return transactionSaved, nil
}

func (s *service) GetAllTransaction(params ParamsGetAllTransaction, url string) (PaginationTransaction, error) {
	var paginationTransaction PaginationTransaction
	offset := params.Page * params.Limit

	_, _, _, err := s.transactionRepo.FindAll(params, offset)
	if err != nil {
		return PaginationTransaction{}, err
	}

	paginationTransaction.FirstPage = fmt.Sprintf("%s", "lupa")

	return PaginationTransaction{}, nil

}

func (s *service) UpdateTransaction(input InputEditTransaction, transactionID int) (Transaction, error) {
	// is transaction available
	transactionByID, err := s.transactionRepo.FindByID(transactionID)
	if err != nil {
		return transactionByID, err
	}

	if transactionByID.ID == 0 {
		return transactionByID, fmt.Errorf(
			"transaction not found",
		)
	}

	// binding
	transactionByID.Amount = input.Amount
	transactionByID.Notes = input.Notes
	transactionByID.Type = input.Type

	date, err := helper.TimeParser(input.Date)
	if err != nil {
		return transactionByID, err
	}
	transactionByID.Date = date

	// update
	transactionUpdated, err := s.transactionRepo.Update(transactionByID)
	if err != nil {
		return transactionUpdated, err
	}

	return transactionUpdated, nil
}

func (s *service) DeleteTransactionByID(transactionID int) error {
	// is transaction available
	transactionByID, err := s.transactionRepo.FindByID(transactionID)
	if err != nil {
		return err
	}

	if transactionByID.ID == 0 {
		return fmt.Errorf(
			"transaction not found",
		)
	}

	// delete
	if err := s.transactionRepo.DeleteByID(transactionID); err != nil {
		return err
	}

	return nil
}
