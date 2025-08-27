package main

import "fmt"

func intersection(a, b []int) []int {

	m := make(map[int]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}

	var result []int

	for _, item := range b {
		if _, ok := m[item]; ok {
			result = append(result, item)
			delete(m, item)
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3, 5, 7, 7, 10}
	B := []int{2, 2, 4, 7, 8, 10}

	C := intersection(A, B)

	fmt.Printf("Слайс A: %v\n", A)
	fmt.Printf("Слайс B: %v\n", B)
	fmt.Printf("Пересечение: %v\n", C) // Вывод: Пересечение: [2 3 8]
}
