package isogram

import (
	"strings"
	"unicode"
)

const nLetters = 26

//IsIsogram verify if a given word is an isogram
func IsIsogram(word string) bool {
	seen := make(map[rune]bool, nLetters)
	for _, l := range strings.ToLower(word) {
		if !unicode.IsLetter(l) {
			continue
		}
		if seen[l] {
			return false
		}
		seen[l] = true
	}
	return true
}
