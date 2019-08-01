package grains

import (
	"errors"
)

//Square returns the number of grains on the square
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("invalid square")
	}

	sum := uint64(1) << uint64(square-1)

	return sum, nil
}

//Total calculates the total of grains in the board
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		sum, _ := Square(i)
		total += sum
	}
	return total
}
