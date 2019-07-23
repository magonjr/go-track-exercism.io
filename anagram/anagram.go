package anagram

import (
	"strings"
)

//Detect verify cadidates returning only the anagrams of a given word
func Detect(word string, candidates []string) []string {
	anagrams := make([]string, 0, len(candidates))
	word = strings.ToLower(word)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)

		if len(candidate) != len(word) || word == candidateLower {
			continue
		}

		temp := []rune(candidateLower)
		for _, l := range word {
			for i := 0; i < len(temp); i++ {
				if temp[i] == l {
					temp[i] = ' '
					break
				}
			}
		}

		if strings.TrimSpace(string(temp)) == "" {
			anagrams = append(anagrams, candidate)
		}

	}

	return anagrams
}
