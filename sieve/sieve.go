package sieve

//Sieve returns a slice with the primes in the n numbers
func Sieve(n int) []int {
	multiple := make([]bool, n)
	primes := make([]int, 0, n)

	if n > 0 {
		multiple[0] = true
	}

	for i := 2; i <= n; i++ {
		for j := i + i; j <= n; j += i {
			multiple[j-1] = true
		}
	}

	for i := 0; i < n; i++ {
		if !multiple[i] {
			primes = append(primes, i+1)
		}
	}
	return primes
}
