package helper

import (
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

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
