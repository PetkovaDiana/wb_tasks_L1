package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Rectangle представляет старый класс прямоугольника
type Rectangle struct {
	Width  int
	Height int
}

// Старый метод для получения площади прямоугольника
func (r *Rectangle) GetArea() int {
	return r.Width * r.Height
}

// Shape интерфейс для новой системы
type Shape interface {
	Area() int
}

// RectangleAdapter адаптер для преобразования Rectangle в Shape
type RectangleAdapter struct {
	Rect *Rectangle
}

// Реализация метода Area для интерфейса Shape
func (adapter *RectangleAdapter) Area() int {
	return adapter.Rect.GetArea()
}

func main() {
	// Создаем старый объект Rectangle
	rect := &Rectangle{
		Width:  10,
		Height: 5,
	}

	// Используем адаптер для работы с интерфейсом Shape
	shape := &RectangleAdapter{
		Rect: rect,
	}

	// Выводим площадь через интерфейс Shape
	fmt.Printf("Rectangle area: %d\n", shape.Area())
}
