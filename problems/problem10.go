package problems

import (
	"github.com/kaezarrex/project-euler/util"
)

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/
func Euler10(limit int) (answer int) {
	c := make(chan int)
	go util.Primes(c)

	for i := range c {
		if i >= limit {
			break
		}
		answer += i
	}
	return answer
}
