package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	s1 := "главрыба"
	fmt.Println("Оригинал 1: ", s1)

	reversed1 := reverseString(s1)
	fmt.Println("Результат 1: ", reversed1)

	s2 := "Hello 🌍"
	fmt.Println("Оригинал 2: ", s2)

	reversed2 := reverseString(s2)
	fmt.Println("Результат 2: ", reversed2)
}
