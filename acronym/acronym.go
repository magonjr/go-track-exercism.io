package acronym

import "strings"

// Abbreviate generates an acronym
func Abbreviate(s string) string {
	var acr string
	init := true
	for _, x := range []rune(s) {
		if x == ' ' || x == '-' || x == '_' {
			init = true
			continue
		}
		if init {
			acr += strings.ToUpper(string(x))
			init = false
		}
	}
	return acr
}
