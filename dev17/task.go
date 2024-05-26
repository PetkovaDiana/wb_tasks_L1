package main

import (
	"fmt"
)

/*Реализовать бинарный поиск встроенными методами языка.*/

func main() {
	// Создаем отсортированный срез
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Элемент, который мы ищем
	target := 5

	// Выполняем бинарный поиск
	index := binarySearch(data, target)

	// Выводим результат
	if index != -1 {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}

// Бинарный поиск
func binarySearch(data []int, target int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid] == target {
			return mid // Элемент найден, возвращаем его индекс
		} else if data[mid] < target {
			low = mid + 1 // Искать справа от середины
		} else {
			high = mid - 1 // Искать слева от середины
		}
	}

	return -1 // Элемент не найден
}
