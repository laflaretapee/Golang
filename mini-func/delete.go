package main

import (
	"fmt"
)

func delete(a []int) []int {

	mapa := make(map[int]bool)

	res := []int{}

	for _, v := range a {
		if !mapa[v] {
			res = append(res, v)
		}
		mapa[v] = true
	}

	return res

}

func main() {
	sl := []int{1, 2, 2, 2, 2, 3, 4, 5, 6, 6, 6, 8}

	fmt.Println(delete(sl))
}
