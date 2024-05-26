package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {

	a := 10
	b := 5

	a, b = b, a

	fmt.Printf("a = %d, b = %d\n", a, b)

}
