package main

import (
	"fmt"
)

func revers(a string) string {

	res := make([]rune, 0, len(a))

	runes := []rune(a)

	for i := len(runes) - 1; i >= 0; i-- {
		res = append(res, runes[i])
	}

	return string(res)
}

func main() {
	aStr := "Dinar"

	fmt.Println(revers(aStr))
}
