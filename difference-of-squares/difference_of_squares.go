package diffsquares

//SumOfSquares calculates the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

//SquareOfSum calculates the  square of the sum of  the first n natural numbers.
func SquareOfSum(n int) int {
	sum := (n * (1 + n)) / 2
	return sum * sum
}

//Difference calculates the difference between the square of the sum and the sum of the squares of the first n natural numbers.
func Difference(n int) int {
	diference := SquareOfSum(n) - SumOfSquares(n)

	return diference
}
