package main

import "fmt"

func main() {
	squreSum := 0
	sumSqure := 0

	for i := 1; i <= 100; i++ {
		squreSum += i * i
		sumSqure += i
	}

	sumSqure *= sumSqure

	result := sumSqure - squreSum

	fmt.Println(result)
}
