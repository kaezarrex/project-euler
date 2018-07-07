package problems

import (
	"math"
)

func isPrime(x int) bool {
	for i := 2; i < x/2+1; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
func Euler3(limit int) (answer int) {
	sqrt := int(math.Sqrt(float64(limit)))

	for i := sqrt; i > 1; i-- {
		if limit%i == 0 {
			if isPrime(limit / i) {
				return limit / i
			}
			if isPrime(i) {
				return i
			}
		}
	}

	return 1
}
