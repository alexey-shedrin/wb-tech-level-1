package main

import "fmt"

func main() {
	a := 3
	b := 7

	fmt.Println("Обмен значений арифметическим способом:")
	fmt.Printf("Исходные значения: a = %d, b = %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

	c := 3
	d := 7

	fmt.Println("Обмен значений через XOR:")
	fmt.Printf("Исходные значения: c = %d, d = %d\n", c, d)

	c = c ^ d
	d = c ^ d
	c = c ^ d

	fmt.Printf("После обмена: c = %d, d = %d\n", c, d)
}
