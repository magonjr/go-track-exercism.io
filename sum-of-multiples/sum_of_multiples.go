package summultiples

//SumMultiples find the sum of all the unique multiples of particular numbers up to but not including that number.
func SumMultiples(n int, values ...int) int {
	var sum int
	for i := 1; i < n; i++ {
		for _, k := range values {
			if k <= 0 {
				break
			}

			if i%k == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
