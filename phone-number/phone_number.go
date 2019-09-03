package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

//Number validates and remove nondigit caracteres from a given fone number
func Number(s string) (string, error) {
	digits := ""
	for _, d := range s {
		if unicode.IsDigit(d) {
			digits = digits + string(d)
		}
	}

	if len(digits) < 10 {
		return "", errors.New("invalid digits length")
	}

	if len(digits) > 10 {
		if digits[0] != '1' {
			return "", errors.New("invalid start")
		}
		digits = digits[1:]
	}

	if digits[0] < '2' || digits[3] < '2' {
		return "", errors.New("invalid digit")
	}

	return digits, nil
}

//AreaCode provides de area code from a given fone number
func AreaCode(s string) (string, error) {
	s, err := Number(s)
	if err != nil {
		return "", err
	}
	return s[:3], nil
}

//Format generates a fone number in the format (NXX)-NXX-XXXX
func Format(s string) (string, error) {
	s, err := Number(s)
	if err != nil {
		return "", err
	}
	s = fmt.Sprintf("(%s) %s-%s", s[:3], s[3:6], s[6:])
	return s, nil
}
