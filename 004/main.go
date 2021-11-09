package main

import (
	"fmt"
	"strconv"
)

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

func isPalindromic(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	output := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			result := i * j
			if isPalindromic(strconv.Itoa(result)) {
				if output <= result {
					output = result
				}
			}
		}
	}
	fmt.Println(output)

}
