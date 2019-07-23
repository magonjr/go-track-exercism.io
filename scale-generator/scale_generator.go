package scale

import (
	"strings"
)

var scale = [2][12]string{
	{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"},
	{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"},
}

var scaleMap = map[string]int{
	"C": 0,
	"a": 0,

	"G": 0, "D": 0, "A": 0, "E": 0, "B": 0, "F#": 0,
	"e": 0, "b": 0, "f#": 0, "c#": 0, "g#": 0, "d#": 0,

	"F": 1, "Bb": 1, "Eb": 1, "Ab": 1, "Db": 1, "Gb": 1,
	"d": 1, "g": 1, "c": 1, "f": 1, "bb": 1, "eb": 1}

//Scale generate a scale from a tonic and an interval
func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}
	newScale := make([]string, 0, len(interval))
	scaleIndex := scaleMap[tonic]
	tonic = verifyMinors(tonic)
	var i int
	for i = 0; i < 12; i++ {
		if scale[scaleIndex][i] == tonic {
			break
		}
	}

	for _, m := range []rune(interval) {
		newScale = append(newScale, scale[scaleIndex][i])
		i++
		if m == 'M' {
			i++
		}
		if m == 'A' {
			i += 2
		}
		i = i % 12
	}
	return newScale
}

func verifyMinors(tonic string) string {
	if len(tonic) == 1 {
		return strings.ToUpper(tonic)
	}
	return strings.ToUpper(tonic[0:1]) + tonic[1:]
}
