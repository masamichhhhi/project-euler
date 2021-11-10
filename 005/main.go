package main

import "fmt"

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func divisibleTo20(input int) bool {
	for i := 1; i <= 20; i++ {
		if input%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	target := 20
	for !divisibleTo20(target) {
		target++
	}
	fmt.Println(target)
}
