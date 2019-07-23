package isbn

import (
	"unicode"
)

//IsValidISBN verify an isbn string retuning whether or not it is valid
func IsValidISBN(s string) bool {
	digits := 0
	sum := 0

	for i, r := range s {
		if r == '-' {
			continue
		}

		if !unicode.IsDigit(r) && (i != len(s)-1 || r != 'X') {
			return false
		}

		digits++

		if digits > 10 {
			return false
		}

		if r == 'X' {
			sum += 10
			break
		}

		sum += (11 - digits) * int(r-'0')
	}

	if digits != 10 {
		return false
	}

	return sum%11 == 0
}
