package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Засыпаем...")

	duration := 3 * time.Second
	sleep(duration)

	fmt.Println("Проснулись!")
}
