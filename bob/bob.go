//Package bob contains bob exercism exercice
package bob

import (
	"strings"
	"unicode"
)

/* Hey func
Bob answers 'Sure.' if you ask him a question.
He answers 'Whoa, chill out!' if you yell at him.
He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
He says 'Fine. Be that way!' if you address him without actually saying anything.
He answers 'Whatever.' to anything else.
*/
func Hey(remark string) string {
	text := []rune(strings.TrimSpace(remark))
	question := false
	yell := true
	letters := false

	for i := 0; i < len(text); i++ {
		if !unicode.IsLetter(text[i]) {
			continue
		}

		letters = true

		if unicode.IsLower(text[i]) {
			yell = false
			break
		}
	}

	if !letters {
		yell = false
	}

	if len(text) > 0 && text[len(text)-1] == '?' {
		question = true
	}

	switch {
	case question && yell:
		return "Calm down, I know what I'm doing!"
	case question:
		return "Sure."
	case yell:
		return "Whoa, chill out!"
	case len(text) == 0:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}
