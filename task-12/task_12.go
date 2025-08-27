package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})

	for _, item := range sequence {
		set[item] = struct{}{}
	}

	uniqueSequence := make([]string, 0, len(set))
	for key := range set {
		uniqueSequence = append(uniqueSequence, key)
	}

	fmt.Println("Исходная последовательность:", sequence)
	fmt.Println("Уникальная последовательность:", uniqueSequence)
}
