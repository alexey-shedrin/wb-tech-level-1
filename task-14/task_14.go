package main

import "fmt"

func outputType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Переменная имеет тип: int")
	case string:
		fmt.Println("Переменная имеет тип: string")
	case bool:
		fmt.Println("Переменная имеет тип: bool")
	case chan int:
		fmt.Println("Переменная имеет тип: chan int")
	case chan string:
		fmt.Println("Переменная имеет тип: chan string")
	case chan bool:
		fmt.Println("Переменная имеет тип: chan bool")
	default:
		fmt.Printf("Переменная имеет другой тип: %T\n", v)
	}
}

func main() {
	var a int = 1
	var b string = "a"
	var c bool = true
	var d chan int = make(chan int)
	var e chan string = make(chan string)
	var f chan bool = make(chan bool)
	var g float64 = 1

	outputType(a)
	outputType(b)
	outputType(c)
	outputType(d)
	outputType(e)
	outputType(f)
	outputType(g)
}
