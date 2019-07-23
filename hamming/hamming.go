package hamming

import "errors"

//Distance calculates hhaming distancve beetween a and b DNA sequences
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("invalid lengths")
	}
	var distance int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
