package romannumerals

import (
	"errors"
	"strings"
)

//ToRomanNumeral converts to roman
func ToRomanNumeral(arabic int) (string, error) {
	var roman string

	if arabic < 1 || arabic > 3000 {
		return "", errors.New("invalid number")
	}
	//I, V, X, L, C, D, M
	if arabic > 999 {
		roman += strings.Repeat("M", arabic/1000)
		arabic = arabic % 1000
	}
	if arabic > 899 {
		roman += "CM"
		arabic = arabic % 900
	}
	if arabic > 499 {
		roman += strings.Repeat("D", arabic/500)
		arabic = arabic % 500
	}
	if arabic > 399 {
		roman += "CD"
		arabic = arabic % 400
	}
	if arabic > 99 {
		roman += strings.Repeat("C", arabic/100)
		arabic = arabic % 100
	}
	if arabic > 89 {
		roman += "XC"
		arabic = arabic % 90
	}
	if arabic > 49 {
		roman += strings.Repeat("L", arabic/50)
		arabic = arabic % 50
	}
	if arabic > 30 {
		roman += "XL"
		arabic = arabic % 40
	}
	if arabic > 10 {
		roman += strings.Repeat("X", arabic/10)
		arabic = arabic % 10
	}
	if arabic > 8 {
		roman += "IX"
		arabic = 0
	}
	if arabic > 4 {
		roman += "V"
		arabic = arabic % 5
	}
	if arabic > 3 {
		roman += "IV"
		arabic = 0
	}
	if arabic > 0 {
		roman += strings.Repeat("I", arabic)
	}

	return roman, nil

}
