package main

import (
	"fmt"
)

func sum(a int) int {
	if a == 1 {
		return 1
	}

	return a + sum(a-1)

}

func main() {
	x := 4

	fmt.Println(sum(x))
}
