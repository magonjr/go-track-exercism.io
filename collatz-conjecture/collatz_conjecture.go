package collatzconjecture

import "errors"

//CollatzConjecture calculates the CollatzConjecture
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid input")
	}

	steps := 0

	for n != 1 {
		steps++
		//If n is even, divide n by 2 to get n / 2. If n is odd, multiply n by 3 and add 1 to get 3n + 1.
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}

	return steps, nil
}
