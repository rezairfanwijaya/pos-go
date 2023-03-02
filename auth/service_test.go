package auth

import "testing"

var AuthService = authService{}

func TestGenerateToken(t *testing.T) {
	token, err := AuthService.GenerateTokenJWT(12)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)
}

func TestVerifyTokenJWT(t *testing.T) {
	token, err := AuthService.VerifyTokenJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMn0.2Dz18ykezo5Xyoqm81aakzU154jSmnreP9oe00XC4Z4")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)
}
