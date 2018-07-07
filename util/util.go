package util

import (
	"math"
)

// Sum adds all integers in a slice
func Sum(nums []int) int {
	result := 0
	for _, x := range nums {
		result += x
	}
	return result
}

// Fibonacci sends consecutive fibonacci numbers to a channel
func Fibonacci(c chan int) {
	x, y := 0, 1
	for {
		c <- x
		x, y = y, x+y
	}
}

// Primes sends consecutive prime numbers to a channel
func Primes(c chan int) {
	primes := []int{}
PrimeLoop:
	for i := 2; ; i++ {
		sqrt := int(math.Ceil(math.Sqrt(float64(i))))
		for _, p := range primes {
			if i%p == 0 {
				continue PrimeLoop
			}
			if p > sqrt {
				break
			}
		}
		primes = append(primes, i)
		c <- i
	}
}
