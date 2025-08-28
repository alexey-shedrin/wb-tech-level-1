package main

import (
	"fmt"
)

type Order struct {
	UID string
}

func (s *Order) PrintUID() {
	fmt.Println("Order UID:", s.UID)
}

type Outputter interface {
	Print()
}

type OrderAdapter struct {
	UID     string
	Comment string
}

func NewOrderAdapter(order *Order, comment string) *OrderAdapter {
	return &OrderAdapter{
		UID:     order.UID,
		Comment: comment,
	}
}

func (a *OrderAdapter) Print() {
	fmt.Printf("Order UID: %s, Comment: %s\n", a.UID, a.Comment)
}

func PrintOrderInfo(outputter Outputter) {
	outputter.Print()
}

func main() {
	order := &Order{UID: "12345"}

	newOrder := NewOrderAdapter(order, "Не звонить в дверь")

	PrintOrderInfo(newOrder)
}

/*
	Название Адаптер говорит само за себя - такой подход позволяет
	адаптировать код к легаси коду и сторонним библиотекам.
	Плюсы - изолирование клиентского кода от реализации отдельных
	структур и их методов, возможность повторного использования кода,
	улучшение читаемости в main.
	Минусы - усложнение и возможная избыточность кода.
	Пример реального использования - драйверы-адаптеры для различных
	баз данных. Например, github.com/lib/pq для работы с PostgreSQL.
*/
