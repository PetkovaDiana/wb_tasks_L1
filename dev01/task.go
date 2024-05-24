package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Определение структуры Human
type Human struct {
	Name string
	Age  int
}

// Talk - метод для структуру human
func (h *Human) Talk() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", h.Name, h.Age)
}

// Swim - другой метод для структуру human
func (h *Human) Swim() {
	fmt.Printf("I'm swimming\n")
}

// Action - структура, в которую встроили структуру Human,
// для того чтобы наследовать поля и методы структуры Human
type Action struct {
	Human
}

// Do - метод для структуру Action
func (a *Action) Do() {
	fmt.Printf("%s is now swimming", a.Name)
}

func main() {
	// Создаем экземпляра структуры Action
	action := Action{
		Human: Human{
			Name: "Diana",
			Age:  22,
		},
	}

	//Вызов методов структуру Human через экземпляр структуру Action
	action.Talk() // Talk() - вызов метода из структуры Human
	action.Swim() // Swim() - вызов метода из структуры Human
	action.Do()   //Do() - вызов метода из структуры Human
}

// В гошке нет прямого наследования, как в ОПП языках, но мы можем использовать одну конструкцию,
// которая позволяет нам встраивать методы одной структуру в другую.
// В ОПП  - это называется полиморфизмом и достигается с помощью интерфейсов.
// В данном примере структура Action может наследовать методы структуру Human
