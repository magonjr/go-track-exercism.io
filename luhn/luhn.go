package luhn

import (
	"unicode"
)

//Valid Given a number determine whether or not it is valid per the Luhn formula.
func Valid(s string) bool {
	code := []rune(s)
	double := false
	sum := 0
	digits := 0

	for i := len(code) - 1; i >= 0; i-- {
		if !unicode.IsDigit(code[i]) {
			if code[i] != ' ' {
				return false
			}
			continue
		}

		digits++

		digit := int(code[i] - '0')

		if double {
			digit *= 2
		}

		if digit > 9 {
			digit -= 9
		}

		sum += digit

		double = !double
	}

	return digits > 1 && sum%10 == 0
}
