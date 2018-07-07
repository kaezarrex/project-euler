package problems

import (
	"math"
	"strconv"
)

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	if s[0] == s[len(s)-1] {
		return isPalindrome(s[1 : len(s)-1])
	}
	return false
}

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two
2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
func Euler4(length int) int {
	low, high := int(math.Pow10(length-1)), int(math.Pow10(length)-1)

	for i, j, k, second := high, high, high, false; k >= low; {
		if isPalindrome(strconv.Itoa(i * j)) {
			return i * j
		}
		if i < high {
			i++
			j--
			continue
		}
		if !second {
			i = k
			j = k - 1
			second = true
			continue
		}
		k--
		i = k
		j = k
		second = false
	}
	return 0
}
