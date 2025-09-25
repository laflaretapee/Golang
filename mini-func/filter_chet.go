package main

import (
	"fmt"
)

func filter(a []int) []int {

	res := make([]int, 0, len(a))

	for _, v := range a {
		if v%2 == 0 {
			res = append(res, v)
		}
	}

	return res

}

func main() {
	sl := []int{}

	fmt.Println(filter(sl))
}
