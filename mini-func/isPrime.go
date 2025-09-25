package main

import (
	"fmt"
	"math"
)

func isPrime(a int) bool {

	if a <= 1 {
		return false
	}

	if a == 2 {
		return true
	}

	if a%2 == 0 {
		return false
	}

	aSqrt := int(math.Sqrt(float64(a)))

	for i := 3; i < aSqrt; i += 2 {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	b := 19

	fmt.Println(isPrime(b))
}
