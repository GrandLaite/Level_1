/*
 Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Say() {
	fmt.Printf("Привет, меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

type Action struct {
	Human
	Activity string
}

func (a Action) PersonAction() {
	fmt.Printf("%s сейчас %s.\n", a.Name, a.Activity)
}

func main() {
	person := Action{
		Human: Human{
			Name: "Маша",
			Age:  30,
		},
		Activity: "спит",
	}

	person.Say()

	person.PersonAction()
}
