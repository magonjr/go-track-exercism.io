package atbash

import (
	"unicode"
)

const block = 5

//Atbash econdes a string in the Atbash format
func Atbash(s string) string {

	letters := []rune(s)
	encoded := make([]rune, 0, len(letters)*2)
	count := 0

	for _, r := range letters {

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}

		if count == block {
			encoded = append(encoded, ' ')
			count = 0
		}

		if unicode.IsDigit(r) {
			encoded = append(encoded, r)
			count++
			continue
		}

		reverse := 'a' + ('z' - unicode.ToLower(r))
		encoded = append(encoded, reverse)
		count++
	}

	return string(encoded)
}
