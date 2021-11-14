package main

import "fmt"

// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

// a^2 + b^2 = c2
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 52.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

func isPythagorean(a, b, c int) bool {
	if (a*a + b*b) == c*c {
		return true
	} else {
		return false
	}

}

func isSum1000(a, b, c int) bool {
	if a+b+c == 1000 {
		return true
	} else {
		return false
	}
}

func main() {
	var result int

	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			for c := 1; c < 1000; c++ {
				if isPythagorean(a, b, c) && isSum1000(a, b, c) {
					result = a * b * c
					break
				}
			}
		}

	}

	fmt.Println(result)
}
