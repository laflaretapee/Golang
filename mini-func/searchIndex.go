package main

import (
	"fmt"
)

func search(a []int, b int) int {

	for i, v := range a {
		if b == v {
			return i
		}
	}
	return -1
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	ch := 1

	fmt.Println(search(sl, ch))
}
