package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnection(t *testing.T) {
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
				_, err := NewConnection(testCase.Path)
				assert.NotNil(t, err)
			} else {
				db, err := NewConnection(testCase.Path)
				assert.NotNil(t, db)
				assert.Nil(t, err)
				t.Log(db)
			}
		})
	}
}
