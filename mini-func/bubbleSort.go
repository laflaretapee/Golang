package main

import (
	"fmt"
)

func bubbleSort(a []int) []int {

	for i, _ := range a {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {

	as := []int{32, 4231, 543, 76, 98, 7423, 12}

	fmt.Println(bubbleSort(as))

}
