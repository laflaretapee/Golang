package main

import (
	"fmt"
)

func palindrom(a string) bool {
	res := make([]rune, 0, len(a))

	s := []rune(a)

	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}

	if a == string(res) {
		return true
	}
	return false
}

func main() {
	s := "alla"

	fmt.Println(palindrom(s))
}
