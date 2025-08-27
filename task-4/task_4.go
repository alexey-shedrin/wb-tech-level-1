package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Как работает: отдельная горутина ждет сигнал прерывания (нажатия Ctrl+C)
// Как только сигнал получен, вызывается закрытие контекста,
// что приводит к завершению всех воркеров. Также разблокируется
// главная горутина, которая ждет завершения воркеров.

// Контекст будет более предпочтительным, так как ему не нужно знать
// количество воркеров, в отличие от реализации с каналами.

// Доработаем код из предыдущего задания
// Воркер теперь принимает context.Context
func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Используем select, чтобы ждать либо новую задачу, либо сигнал отмены
		select {
		case j, ok := <-jobs:
			if !ok {
				fmt.Printf("Воркер %d: канал задач закрыт, завершаю работу.\n", id)
				return
			}
			fmt.Printf("Воркер %d получил задачу: %d\n", id, j)
			time.Sleep(4 * time.Second)
			fmt.Printf("Воркер %d завершил задачу: %d\n", id, j)
		case <-ctx.Done():
			fmt.Printf("Воркер %d: получен сигнал отмены, завершаю работу.\n", id)
			return
		}
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

	ctx, cancel := context.WithCancel(context.Background())

	// В sigChan записывается сигнал прерывания
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// Обработка сигнала прерывания в отдельной горутине
	go func() {
		<-sigChan
		fmt.Println("\nПолучен сигнал прерывания. Завершаем работу...")
		cancel()
	}()

	jobs := make(chan int, 10)

	var wg sync.WaitGroup

	fmt.Printf("Запускаем %d воркеров... Нажмите Ctrl+C для выхода.\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, &wg)
	}

loop:
	for j := 1; ; j++ {
		select {
		case jobs <- j:
			fmt.Printf("Главная горутина отправила задачу: %d\n", j)
			time.Sleep(time.Second)
		case <-ctx.Done():
			fmt.Println("Главная горутина: прекращаю отправку задач.")
			break loop
		}
	}

	close(jobs)
	fmt.Println("Все задачи отправлены. Ожидание завершения воркеров...")
	wg.Wait()
	fmt.Println("Все воркеры завершены. Программа выключается.")
}
