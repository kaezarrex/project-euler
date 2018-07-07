package problems

import (
	"github.com/kaezarrex/project-euler/util"
)

/*
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
