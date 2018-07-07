package problems

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func Euler9(limit int) (answer int) {
	for c := 1; ; c++ {
		for b := 1; b < c; b++ {
			for a := 1; a < b; a++ {
				if a*a+b*b == c*c && a+b+c == limit {
					return a * b * c
				}
			}
		}
	}

	return
}
