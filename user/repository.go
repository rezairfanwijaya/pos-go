package user

import "gorm.io/gorm"

type IRepository interface {
	FindByUsername(username string) (User, error)
	FindByUserID(userID int) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUsername(username string) (User, error) {
	var user User

	if err := r.db.Where("username = ? ", username).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUserID(userID int) (User, error) {
	var user User

	if err := r.db.Where("id = ?", userID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
