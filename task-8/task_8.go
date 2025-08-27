package main

import "fmt"

func setBitToOne(n int64, i uint) int64 {
	mask := int64(1 << i)
	return n | mask
}

func setBitToZero(n int64, i uint) int64 {
	mask := int64(1 << i)
	return n &^ mask
}

func main() {
	var number1 int64 = 6  // 110
	var position1 uint = 1 // Второй бит справа (нумерация с 0)

	fmt.Printf("Установка 1-го бита числа %d в 0\n", number1)
	fmt.Printf("Исходное число: %d (%b)\n", number1, number1)

	result1 := setBitToZero(number1, position1)

	fmt.Printf("Результат: %d (%b)\n", result1, result1) // Ожидаем 4 (100)

	var number2 int64 = 5  // 101
	var position2 uint = 1 // Второй бит справа (нумерация с 0)

	fmt.Printf("Установка 1-го бита числа %d в 1\n", number2)
	fmt.Printf("Исходное число: %d (%b)\n", number2, number2)

	result2 := setBitToOne(number2, position2)

	fmt.Printf("Результат: %d (%b)\n", result2, result2) // Ожидаем 7 (111)
}
