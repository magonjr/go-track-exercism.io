package encode

import (
	"strconv"
	"strings"
	"unicode"
)

//RunLengthDecode decode a string in a Run-length encoding (RLE) format
func RunLengthDecode(input string) string {

	var aux, decoded string
	for _, r := range input {

		if unicode.IsDigit(r) {
			aux = aux + string(r)
			continue
		}

		n := 1
		if aux != "" {
			n, _ = strconv.Atoi(aux)
		}
		decoded = decoded + strings.Repeat(string(r), n)
		aux = ""
	}

	return decoded

}

//RunLengthEncode encode a string in a Run-length encoding (RLE) format
func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}

	var (
		encoded string
		count   int
		aux     rune
	)

	for i, r := range input {

		if r == aux || i == 0 {
			aux = r
			count++
			continue
		}

		encoded = addEncode(encoded, aux, count)
		aux = r
		count = 1
	}

	encoded = addEncode(encoded, aux, count)

	return encoded

}

func addEncode(current string, r rune, count int) string {
	if count < 2 {
		return current + string(r)
	}
	return current + strconv.Itoa(count) + string(r)
}
