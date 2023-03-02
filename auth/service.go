package auth

import (
	"fmt"
	"pos/helper"

	"github.com/dgrijalva/jwt-go/v4"
)

type IAuthService interface {
	GenerateTokenJWT(userID int) (string, error)
	VerifyTokenJWT(token string) (*jwt.Token, error)
}

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) GenerateTokenJWT(userID int) (string, error) {
	// payload
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	env, err := helper.GetENV(".env")
	if err != nil {
		return "", err
	}

	// signature
	tokenSigned, err := token.SignedString([]byte(env["SECRET_KEY"]))
	if err != nil {
		return tokenSigned, fmt.Errorf(
			"failed signed token : %v",
			err.Error(),
		)
	}

	return tokenSigned, nil
}

func (s *authService) VerifyTokenJWT(token string) (*jwt.Token, error) {
	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token")
		}

		env, err := helper.GetENV(".env")
		if err != nil {
			return env, err
		}

		return []byte(env["SECRET_KEY"]), nil
	})

	if err != nil {
		return tokenParsed, err
	}

	return tokenParsed, nil
}
