package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	testCases := []struct {
		Name      string
		Path      string
		WantError bool
	}{
		{
			Name:      "success",
			Path:      "../.env",
			WantError: false,
		},
		{
			Name:      "failed",
			Path:      ".env",
			WantError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.WantError {
				_, err := GetENV(testCase.Path)
				assert.NotNil(t, err)
			} else {
				env, err := GetENV(testCase.Path)
				assert.NotNil(t, env)
				assert.Nil(t, err)
				t.Log(env)
			}
		})
	}
}
