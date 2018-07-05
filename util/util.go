package util

func Sum(nums []int) int {
	result := 0
	for _, x := range nums {
		result += x
	}
	return result
}

func Fibonacci(c chan int) {
	x, y := 0, 1
	for {
		c <- x
		x, y = y, x+y
	}
}
