package main

import "fmt"

func deleteElement[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}

	copy(s[i:], s[i+1:])

	s = s[:len(s)-1]

	return s
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 5
	fmt.Printf("Исходный слайс: %v, Длина: %v, Eмкость: %v\n", slice, len(slice), cap(slice))
	fmt.Printf("Удаляем элемент с индексом %v\n", index)

	slice = deleteElement(slice, index)
	fmt.Printf("Итоговый слайс: %v, Длина: %v, Eмкость: %v\n", slice, len(slice), cap(slice))
}
