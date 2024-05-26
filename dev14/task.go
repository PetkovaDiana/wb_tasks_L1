package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	x := 10
	y := "Hello world"
	i := make([]int, 10)

	xType := reflect.TypeOf(x)
	yType := reflect.TypeOf(y)
	iType := reflect.TypeOf(i)

	fmt.Printf("%T: %s\n%T: %s\n%T: %s\n", xType, xType, yType, yType, iType, iType)

}
