package cryptosquare

import (
	"math"
	"unicode"
)

//Encode econdes a string in the square code format
func Encode(s string) string {
	letters := onlyLetters(s)

	var col int = int(math.Sqrt(float64(len(letters))))
	row := col
	nLetters := len(letters)
	for col*row < nLetters {
		col++
		if col*row < nLetters {
			row++
		}
	}

	encoded := make([]rune, 0, len(s))

	for i := 0; i < col; i++ {
		if i > 0 {
			encoded = append(encoded, ' ')
		}
		for j := 0; j < row; j++ {
			letterPos := col*j + i

			if letterPos >= nLetters {
				encoded = append(encoded, ' ')
			} else {
				encoded = append(encoded, letters[letterPos])
			}
		}
	}
	return string(encoded)
}

func onlyLetters(s string) []rune {
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}

		letters = append(letters, unicode.ToLower(r))
	}
	return letters
}
