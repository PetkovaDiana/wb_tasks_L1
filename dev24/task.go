package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

// Point- структура наши точек
type Point struct {
	X float64
	Y float64
}

// NewPoint - конструктор
func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// String - выводит координаты точки
func (p *Point) String() string {
	return fmt.Sprintf("x: %f\ny: %f", p.X, p.Y)
}

// Distance - вычисление расстояния между двумя точками
func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p2.X-p1.X), 2) + math.Pow((p2.Y-p1.Y), 2))
}

func main() {
	p1 := NewPoint(15, 15)
	p2 := NewPoint(30, 13)

	fmt.Println(p1)
	fmt.Println()
	fmt.Println(p2)
	fmt.Println("Distance between two point is: ", Distance(p1, p2))
}
