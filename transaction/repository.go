package transaction

import (
	"math"

	"gorm.io/gorm"
)

type IRepository interface {
	Save(transaction Transaction) (Transaction, error)
	FindByID(transactionID int) (Transaction, error)
	FindAll(params ParamsGetAllTransaction, offset int) ([]Transaction, int, int, error)
	Update(transaction Transaction) (Transaction, error)
	DeleteByID(transactionID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	if err := r.db.Create(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindByID(transactionID int) (Transaction, error) {
	var transaction Transaction
	if err := r.db.Where("id = ?", transactionID).Find(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindAll(params ParamsGetAllTransaction, offset int) ([]Transaction, int, int, error) {
	var transactions []Transaction
	var totalData int64 = 0
	var totalPage int = 0

	if err := r.db.Where("type LIKE ? ", "%"+params.TransactionType+"%").Where("amount BETWEEN ? AND ?", params.FromAmount, params.ToAmount).Order(params.Amount).Order(params.Date).Limit(params.Limit).Offset(offset).Find(&transactions).Error; err != nil {
		return transactions, int(totalData), totalPage, err
	}

	// total data
	if err := r.db.Model(&Transaction{}).Count(&totalData).Error; err != nil {
		return transactions, int(totalData), totalPage, err
	}

	// total page
	totalPage = int(math.Ceil(float64(totalData)/float64(params.Limit))) - 1

	return transactions, int(totalData), totalPage, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	if err := r.db.Save(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) DeleteByID(transactionID int) error {
	var transaction Transaction
	if err := r.db.Where("id = ?", transactionID).Delete(&transaction).Error; err != nil {
		return err
	}

	return nil
}
