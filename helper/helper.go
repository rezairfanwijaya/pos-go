package helper

import (
	"github.com/joho/godotenv"
)

func GetENV(path string) (env map[string]string, err error) {
	env, err = godotenv.Read(path)
	if err != nil {
		return env, err
	}

	return env, nil
}
