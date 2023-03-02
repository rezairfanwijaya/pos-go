package database

import (
	"fmt"
	"pos/helper"
	"pos/user"

	"gorm.io/gorm"
)

func migration(db *gorm.DB) (err error) {
	// pastikan belum ada data di table users
	// untuk melakukan up migration
	empty, err := isEmpty(db)
	if err != nil {
		return err
	}

	if empty {
		password, err := helper.HashPassword("12345")
		if err != nil {
			return fmt.Errorf(
				"failed hash password when migration : %v",
				err.Error(),
			)
		}

		if err := db.Create(&user.User{
			Username: "admin",
			Password: password,
		}).Error; err != nil {
			return fmt.Errorf(
				"failed store data user : %v",
				err.Error(),
			)
		}

		return nil
	}

	return nil
}

func isEmpty(db *gorm.DB) (bool, error) {
	var users []user.User

	if err := db.Find(&users).Error; err != nil {
		return false, err
	}

	if len(users) == 0 {
		return true, nil
	}

	return false, nil
}
