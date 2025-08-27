package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayHello() {
	fmt.Printf("Привет! Меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

func (h *Human) Introduce() {
	fmt.Printf("Я человек по имени %s.\n", h.Name)
}

type Action struct {
	Human
	Activity string
}

func (a *Action) PerformAction() {
	fmt.Printf("%s сейчас выполняет действие: %s.\n", a.Name, a.Activity)
}

func (a *Action) Introduce() {
	fmt.Printf("Я активная сущность, и моё имя %s. Я люблю %s.\n", a.Name, a.Activity)
}

func main() {
	personAction := Action{
		Human: Human{
			Name: "Иван",
			Age:  35,
		},
		Activity: "путешествовать",
	}

	personAction.SayHello()
	personAction.PerformAction()
	personAction.Introduce()
	personAction.Human.Introduce()
}
