package problems

import (
	"github.com/kaezarrex/project-euler/util"
)

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is
13.

What is the 10 001st prime number?
*/
func Euler7(limit int) int {
	var count = 0
	c := make(chan int)
	go util.Primes(c)

	for i := range c {
		count++
		if count >= limit {
			return i
		}
	}
	return 0
}
