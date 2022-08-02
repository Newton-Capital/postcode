package postcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		description string
		country     string
		input       string
		expected    error
	}{
		{
			description: "Paris, France",
			country:     "France",
			input:       "75008",
			expected:    fmt.Errorf("unsupported country"),
		},
		{
			description: "United Kingdom, London",
			country:     "United Kingdom",
			input:       "B24 8DF",
			expected:    nil,
		},
		{
			description: "Calgary, Canada",
			country:     "Canada",
			input:       "T2E 4Y5",
			expected:    nil,
		},
	}

	for _, testCase := range testCases {
		t.Logf("Validating `%s` (%s)", testCase.input, testCase.description)
		err := Validate(testCase.input, testCase.country)
		if err != nil {
			assert.Equal(t, testCase.expected.Error(), err.Error())
		}
	}
}
