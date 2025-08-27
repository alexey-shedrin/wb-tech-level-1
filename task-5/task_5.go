package main

import (
	"fmt"
	"time"
)

func main() {
	var n int

	fmt.Print("Введите количество секунд для таймаута: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	ch := make(chan int)

	go func() {
		for i := 1; ; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	timeoutDuration := time.Duration(n) * time.Second
	timeoutChan := time.After(timeoutDuration)
	fmt.Printf("Программа будет работать %v.\n", timeoutDuration)

	for {
		select {
		case val := <-ch:
			fmt.Printf("Получено значение: %d\n", val)
		case <-timeoutChan:
			fmt.Println("Таймаут! Программа завершается.")
			return
		}
	}
}
