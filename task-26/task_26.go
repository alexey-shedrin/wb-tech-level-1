package main

import (
	"fmt"
	"strings"
)

func checkUniqueSymbols(s string) bool {
	s = strings.ToLower(s)

	symbols := make(map[rune]struct{})

	for _, v := range s {
		if _, ok := symbols[v]; ok {
			return false
		}
		symbols[v] = struct{}{}
	}

	return true
}

func main() {
	fmt.Printf("'abcd' -> %t\n", checkUniqueSymbols("abcd"))
	fmt.Printf("'abCdefAaf' -> %t\n", checkUniqueSymbols("abCdefAaf"))
	fmt.Printf("'aabcd' -> %t\n", checkUniqueSymbols("aabcd"))
}
