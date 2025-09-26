package main

import "fmt"

func chet(a string) int {

	inWord := false

	k := 0

	s := []rune(a)

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			inWord = false
		} else if !inWord {
			k++
			inWord = true
		}

	}

	return k

}

func main() {
	s := "Hello Dinar how are you?"

	fmt.Println(chet(s))
}
