package main

import "fmt"

func min_max(a []int) (max int, min int) {

	min = 1000000

	max = 0

	for _, v := range a {
		if v > max {
			max = v
		}
	}

	for _, v := range a {
		if v < min {
			min = v
		}
	}

	return max, min

}

func main() {

	a := []int{1, 2, 3, 4, 5}

	maxVal, minVAl := min_max(a)

	fmt.Println(maxVal, minVAl)

}
