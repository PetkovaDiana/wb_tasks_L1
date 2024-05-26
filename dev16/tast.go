package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	// Создаем исходный неотсортированный срез
	data := generateSlice(10)

	fmt.Println("Before sorting:", data)

	// Сортируем срез
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	fmt.Println("After sorting:", data)
}

// Генерация неотсортированного среза случайных чисел
func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) // Генерируем случайные числа от 0 до 99
	}
	return slice
}
