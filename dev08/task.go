package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// Функция для установки i-го бита в 1
func setBit(n int64, i uint) int64 {
	return n | (1 << i)
}

// Функция для установки i-го бита в 0
func clearBit(n int64, i uint) int64 {
	return n &^ (1 << i)
}

func main() {
	var num int64 = 42 // Пример числа
	var i uint = 3     // Индекс бита для изменения (нумерация начинается с 0)

	// Установка i-го бита в 1
	num = setBit(num, i)
	fmt.Printf("After setting bit %d to 1: %d\n", i, num)

	// Установка i-го бита в 0
	num = clearBit(num, i)
	fmt.Printf("After clearing bit %d to 0: %d\n", i, num)
}

// 42 = 0 0 1 0 1 0 1 0
//    * 0 0 0 0 1 0 0 0
//      _______________
//      0 0 1 0 0 0 1 0  =  34
