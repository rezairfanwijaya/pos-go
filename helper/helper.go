package helper

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func GenerateResponse(status int, message string, data interface{}) responseAPI {
	return responseAPI{
		Meta: meta{
			Status:  status,
			Message: message,
		},
		Data: data,
	}
}

func GetENV(path string) (env map[string]string, err error) {
	env, err = godotenv.Read(path)
	if err != nil {
		return env, err
	}

	return env, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func VerifyPassword(rawPassword, hashedPassword []byte) (err error) {
	err = bcrypt.CompareHashAndPassword(
		hashedPassword,
		rawPassword,
	)

	if err != nil {
		return err
	}

	return nil
}

func FormatingErrorBinding(err error) (errBinding []string) {
	for _, e := range err.(validator.ValidationErrors) {
		errBinding = append(errBinding, e.Error())
	}

	return errBinding
}

func TimeParser(dateString string) (result time.Time, err error) {
	layout := "2006-01-02 15:04:05"
	result, err = time.Parse(layout, dateString)
	if err != nil {
		return result, fmt.Errorf(
			"failed parsing date: %v",
			err.Error(),
		)
	}

	return result, nil
}
