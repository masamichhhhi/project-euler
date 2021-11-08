package main

import "fmt"

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func main() {
	target := 600851475143
	i := 2
	for i*i <= int(target) {
		if target%i == 0 {
			target = target / i
		} else {
			i++
		}
	}

	fmt.Println(target)
}
