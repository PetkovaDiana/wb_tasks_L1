package main

import (
	"fmt"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	x := 10
	y := "Hello world"
	i := make([]int, 10)

	//xType := reflect.TypeOf(x)
	//yType := reflect.TypeOf(y)
	//iType := reflect.TypeOf(i)

	fmt.Printf("%[1]T: %[1]d\n%[2]T: %[2]s\n%[3]T: %[3]v\n", x, y, i)

}
