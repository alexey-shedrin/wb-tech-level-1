package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func conditionExitWorker(iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Горутина с условным выходом начинает работу.")
	for i := 0; i < iterations; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Условие выполнено, завершаюсь.")
}

func doneChannelWorker(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Горутина с выходом по сигналу из канала начинает работу.")
	<-done
	fmt.Println("Сигнал получен, завершаюсь.")
}

func contextWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Горутина с отменой контекста начинает работу.")
	<-ctx.Done()
	fmt.Println("Контекст отменен, завершаюсь.")
}

func closeChannelWorker(tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Горутина с закрытием канала начинает работу.")
	for range tasks {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Канал закрыт, завершаюсь.")
}

func goexitWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Горутина с вызовом runtime.Goexit() завершена.")
	fmt.Println("Горутина с вызовом runtime.Goexit() начинает работу.")
	runtime.Goexit()
}

func panicWorker(wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Паника обработана")
		}
		wg.Done()
	}()
	fmt.Println("Горутина с паникой начинает работу.")
	panic("паника")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go conditionExitWorker(5, &wg)
	wg.Wait()

	doneChan := make(chan bool)
	wg.Add(1)
	go doneChannelWorker(doneChan, &wg)
	doneChan <- true
	wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go contextWorker(ctx, &wg)
	cancel()
	wg.Wait()

	tasksChan := make(chan int)
	wg.Add(1)
	go closeChannelWorker(tasksChan, &wg)
	tasksChan <- 1
	tasksChan <- 2
	tasksChan <- 3
	close(tasksChan)
	wg.Wait()

	wg.Add(1)
	go goexitWorker(&wg)
	wg.Wait()

	wg.Add(1)
	go panicWorker(&wg)
	wg.Wait()
}
