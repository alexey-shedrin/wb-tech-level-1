package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]

	var less, equal, greater []int
	for _, item := range arr {
		if item < pivot {
			less = append(less, item)
		} else if item == pivot {
			equal = append(equal, item)
		} else {
			greater = append(greater, item)
		}
	}

	lessSorted := quickSort(less)
	greaterSorted := quickSort(greater)

	return append(append(lessSorted, equal...), greaterSorted...)
}

func main() {
	slice1 := []int{10, 7, 8, 9, 1, 5, 3, 6, 4, 2}
	fmt.Println("Неотсортованный слайс 1:", slice1)

	sortedSlice1 := quickSort(slice1)
	fmt.Println("Отсортированный слайс 1:", sortedSlice1)

	slice2 := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println("Неотсортованный слайс 2:", slice2)

	sortedSlice2 := quickSort(slice2)
	fmt.Println("Отсортованный слайс 2:", sortedSlice2)
}
