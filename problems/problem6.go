package problems

/*
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square
of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the
square of the sum.
*/
func Euler6(limit int) int {
	var sumOfSquares, squareOfSum int
	for i := 1; i <= limit; i++ {
		sumOfSquares += i * i
	}
	for i := 1; i <= limit; i++ {
		squareOfSum += i
	}
	squareOfSum *= squareOfSum
	return squareOfSum - sumOfSquares
}
