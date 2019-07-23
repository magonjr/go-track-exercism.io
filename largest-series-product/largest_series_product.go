package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct calculate the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 || span > len(digits) {
		return -1, errors.New("invalid span")
	}

	var maxProd int64

	for i := 0; i <= len(digits)-span; i++ {
		prod, err := prodSequence(digits[i : i+span])

		if err != nil {
			return -1, err
		}

		if prod > maxProd {
			maxProd = prod
		}
	}

	return maxProd, nil
}

func prodSequence(seq string) (int64, error) {
	var prod int64 = 1
	for _, r := range seq {
		if !unicode.IsDigit(r) {
			return 0, errors.New("invalid digit")
		}
		prod *= int64(r - '0')
	}
	return prod, nil
}
