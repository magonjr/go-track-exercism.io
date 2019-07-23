package pangram

import (
	"strings"
	"unicode"
)

const nLetters = 26

//IsPangram determines if a given word is a pangram
func IsPangram(word string) bool {
	seen := make(map[rune]bool, nLetters)
	for _, l := range strings.ToLower(word) {
		if !unicode.IsLetter(l) {
			continue
		}
		if seen[l] {
			continue
		}
		seen[l] = true
	}
	return len(seen) == nLetters
}
