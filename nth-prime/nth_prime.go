package prime

var primes []int

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	if primes == nil {
		primes = make([]int, 4, 1000)
		primes = []int{2, 3, 5, 7}
	}

	count := len(primes)
	x := primes[len(primes)-1]
	for count < n {
		x++
		if isPrime(x) {
			count++
			primes = append(primes, x)
		}
	}
	return primes[n-1], true
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
