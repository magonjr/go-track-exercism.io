package wordcount

import (
	"strings"
	"unicode"
)

//Frequency keeps count ov each word
type Frequency map[string]int

func validLetter(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '\''
}

//WordCount count words in a given sentense
func WordCount(sentence string) Frequency {
	counter := make(map[string]int, 100)
	aux := []rune(strings.ToLower(sentence))

	for k, l := range aux {
		if !validLetter(l) || l == '\t' || l == '\n' {
			aux[k] = ' '
		}
	}
	var i int
	for j, l := range aux {
		if !validLetter(l) {
			word := strings.TrimFunc(string(aux[i:j]), trimSpaceAndCotes)
			i = j + 1
			if word != "" {
				counter[word] = counter[word] + 1
			}
		}
	}

	word := strings.TrimFunc(sentence[i:], trimSpaceAndCotes)
	if word != "" {
		counter[word] = counter[word] + 1
	}
	return counter
}

func trimSpaceAndCotes(r rune) bool {
	return r == ' ' || r == '\''
}
