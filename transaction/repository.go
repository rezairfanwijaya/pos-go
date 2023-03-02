package transaction

import "gorm.io/gorm"

type IRepository interface {
	Save(transaction Transaction) (Transaction, error)
	FindByID(transactionID int) (Transaction, error)
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

func (r *repository) DeleteByID(transactionID int) error {
	var transaction Transaction
	if err := r.db.Where("id = ?", transactionID).Delete(&transaction).Error; err != nil {
		return err
	}

	return nil
}
