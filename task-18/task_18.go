package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	counter := Counter{}

	var wg sync.WaitGroup

	goroutines := 1000
	incrementsForGoroutine := 1000

	fmt.Println("Ожидаемое значение ", goroutines*incrementsForGoroutine)

	for i := 0; i < goroutines; i++ {
		wg.Go(func() {
			for j := 0; j < incrementsForGoroutine; j++ {
				counter.Increment()
			}
		})
	}

	wg.Wait()

	fmt.Println("Итоговое значение ", counter.value)
}
