package main

import "fmt"

func binarySearch(slice []int, target int) int {
	left := 0
	right := len(slice) - 1

	for left <= right {
		middle := (left + right) / 2

		if slice[middle] == target {
			return middle
		}

		if slice[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}

func main() {
	sortedSlice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target1 := 11

	index1 := binarySearch(sortedSlice1, target1)
	fmt.Println("Индекс искомого элемента: ", index1)

	sortedSlice2 := []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283}
	target2 := 44

	index2 := binarySearch(sortedSlice2, target2)
	fmt.Println("Индекс искомого элемента: ", index2)
}
