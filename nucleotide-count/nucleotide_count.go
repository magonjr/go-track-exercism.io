package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

//Counts func that generates Histogram from a
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for i := 0; i < len(d); i++ {
		if val, ok := h[rune(d[i])]; ok {
			h[rune(d[i])] = val + 1
		} else {
			return nil, errors.New("invalid")
		}
	}

	return h, nil
}
