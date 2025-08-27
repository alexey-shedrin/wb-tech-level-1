package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		fmt.Printf("Воркер %d получил задачу: %d\n", id, j)
		time.Sleep(5 * time.Second)
		fmt.Printf("Воркер %d завершил задачу: %d\n", id, j)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров в качестве аргумента.")
		fmt.Println("Пример: go run main.go 5")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Количество воркеров должно быть положительным числом.")
		return
	}

	jobs := make(chan int, 10)

	var wg sync.WaitGroup

	fmt.Printf("Запускаем %d воркеров...\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for j := 1; j <= 30; j++ {
		jobs <- j
		fmt.Printf("Главная горутина отправила задачу: %d\n", j)
		time.Sleep(time.Second)
	}

	close(jobs)
	wg.Wait()
}
