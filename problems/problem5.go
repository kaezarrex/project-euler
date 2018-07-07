package problems

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any
remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func Euler5(limit int) int {
MainLoop:
	for guess := limit * 2; ; guess++ {
		for divisor := limit; divisor > 1; divisor-- {
			if guess%divisor != 0 {
				continue MainLoop
			}
		}
		return guess
	}
}
