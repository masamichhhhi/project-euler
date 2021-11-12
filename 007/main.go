package main

import "fmt"

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10 001st prime number?

func isDividableByPrimes(primes []int, target int) bool {
	for _, prime := range primes {
		if target%prime == 0 {
			return true
		}
	}

	return false
}

func main() {
	target := 14
	primes := []int{2, 3, 5, 7, 11, 13}
	for len(primes) < 10001 {
		if !isDividableByPrimes(primes, target) {
			primes = append(primes, target)
		}

		target++
	}
	fmt.Println(primes[len(primes)-1])
}
