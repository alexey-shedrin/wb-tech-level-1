package main

import "fmt"

func reverseWords(s string) string {
	runes := []rune(s)

	reverseRuneSlice(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			reverseRuneSlice(runes, start, i-1)
			start = i + 1
		}
	}

	reverseRuneSlice(runes, start, len(runes)-1)

	return string(runes)
}

func reverseRuneSlice(r []rune, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func main() {
	s1 := "главрыба больше не глав"
	fmt.Println("Оригинал 1: ", s1)

	reversed1 := reverseWords(s1)
	fmt.Println("Результат 1: ", reversed1)

	s2 := "Hello 🌍"
	fmt.Println("Оригинал 2: ", s2)

	reversed2 := reverseWords(s2)
	fmt.Println("Результат 2: ", reversed2)
}
