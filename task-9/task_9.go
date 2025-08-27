package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 8, 9, 10}

	chX := make(chan int)
	ch2X := make(chan int)

	fmt.Println("Конвейер начал работу...")

	var wg sync.WaitGroup

	wg.Go(func() {
		for _, x := range nums {
			chX <- x
		}
		close(chX)
	})

	wg.Go(func() {
		for x := range chX {
			ch2X <- x * 2
		}
		close(ch2X)
	})

	for result := range ch2X {
		fmt.Println(result)
	}

	wg.Wait()

	fmt.Println("Конвейер завершил работу.")
}
