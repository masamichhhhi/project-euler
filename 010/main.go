package main

import "fmt"

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

func isDividableByPrimes(primes []int, target int) bool {
	for _, prime := range primes {
		if target%prime == 0 {
			return true
		}
	}

	return false
}

func main() {
	limit := 2000000
	primes := []int{2}
	result := 2

	for i := 3; i < limit; i++ {
		if !isDividableByPrimes(primes, i) {
			primes = append(primes, i)
			result += i
		}
	}

	fmt.Println(result)
}
