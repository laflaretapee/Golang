package main

import (
	"fmt"
)

func proiz(a []int, b int) []int {

	res := []int{}

	for _, v := range a {
		res = append(res, v*b)
	}
	return res
}

func main() {
	sl := []int{1, 2, 3, 4, 5}

	fmt.Println(proiz(sl, 2))
}
