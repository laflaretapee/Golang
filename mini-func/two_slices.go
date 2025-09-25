package main

import (
	"fmt"
)

func two_slices(a []int, b []int) []int {
	res := []int{}

	prov := make(map[int]bool)

	for _, v := range append(a, b...) {
		if !prov[v] {
			res = append(res, v)
		}
		prov[v] = true
	}

	return res

}

func main() {
	a := []int{4, 4, 4, 4}

	b := []int{4, 5, 6}

	res := two_slices(a, b)

	fmt.Println(res)
}
