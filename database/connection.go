package database

import (
	"fmt"
	"pos/helper"
	"pos/transaction"
	"pos/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(path string) (*gorm.DB, error) {
	env, err := helper.GetENV(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed get env : %v",
			err.Error(),
		)
	}

	dbUsername := env["DATABASE_USERNAME"]
	dbPassword := env["DATABASE_PASSWORD"]
	dbHost := env["DATABASE_HOST"]
	dbPort := env["DATABASE_PORT"]
	dbName := env["DATABASE_NAME"]

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf(
			"failed create connection : %v",
			err.Error(),
		)
	}

	// migartion schema
	if err := db.AutoMigrate(
		&user.User{},
		&transaction.Transaction{},
	); err != nil {
		return db, fmt.Errorf(
			"failed migration scheme : %v",
			err.Error(),
		)
	}

	if err := migration(db); err != nil {
		return db, fmt.Errorf(
			"failed migration data : %v",
			err.Error(),
		)
	}

	return db, nil
}
