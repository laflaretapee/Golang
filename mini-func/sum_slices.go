package main

import "fmt"

func sum(a []int) int {
	summa := 0

	for _, v := range a {
		summa += v

	}

	return summa
}

func main() {

	a := []int{1, 2, 3, 4, 5}

	s := sum(a)

	fmt.Println(s)
}
