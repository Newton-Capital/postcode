/*
Package postcode validates postal codes. While the validation process does
not guarantee that the postcode actually exists, it does guarantee that the
format of the provided input is valid.

	For more information see:
	https://en.wikipedia.org/wiki/List_of_postal_codes
	https://en.wikipedia.org/wiki/ISO_3166-1

Example

	if err := postcode.Validate("10007"); err != nil {
		// the postcode is not valid
		// treat error
	}
*/
package postcode

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Possible validation errors.
var (
	ErrEmpty          = errors.New("postal code cannot be empty")
	ErrShort          = errors.New("postal code cannot be shorter than 2 characters")
	ErrInvalidCountry = errors.New("invalid country code")
	ErrInvalidFormat  = errors.New("invalid postal code format")
)

// Validate checks if the provided input string matches any of the
// accepted postcode formats. If the validation fails, the function returns
// an error specifying the cause.
func Validate(code string, country string) error {
	if code = strings.ToUpper(strings.TrimSpace(code)); code == "" {
		return ErrEmpty
	}

	format := []rune(code)
	if len(format) < 2 {
		return ErrShort
	}

	// Map input postal code to format.
	countryCode := string(format[:2])
	for i, r := range format {
		switch {
		case unicode.IsDigit(r):
			r = 'N'
		case unicode.IsLetter(r):
			r = 'A'
		}

		format[i] = r
	}

	foundPostalByCountry := ""

	if _, ok := formats[string(format)]; ok {

		// get postal codes format by country
		v, ok := countries[country]
		if ok {

			// loop through the format results to see if any format match the code passed
			// and assign the value to another variable "foundPostalByCountry"
			for _, s := range v {
				if strings.EqualFold(s, string(format)) {
					foundPostalByCountry = s
				}
			}

			// check if variable "foundPostalByCountry" is not empty
			// and return nil,
			// else return error invalid postal code for the country passed
			if !strings.EqualFold(foundPostalByCountry, "") {
				return nil
			} else {
				return fmt.Errorf("invalid postal code for %s", country)
			}

		} else {
			return fmt.Errorf("unsupported country")
		}
	}

	// Check if postal code is valid when accounting for the country code.
	if format[0] == 'A' && format[1] == 'A' {
		format[0], format[1] = 'C', 'C'
		if _, ok := formats[string(format)]; ok {
			if _, ok := countryCodes[countryCode]; !ok {
				return ErrInvalidCountry
			}

			return nil
		}
	}

	return ErrInvalidFormat
}
