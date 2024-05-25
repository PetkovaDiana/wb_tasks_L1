package main

import (
	"fmt"
	"sync"
)

/*
Написать программу,
которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
и выведет их квадраты в stdout.
*/

// Функция для вычисления квадрата числа и отправки результата в канал,
// канал только для записи
func squares(wg *sync.WaitGroup, numbers []int) <-chan int {
	results := make(chan int, len(numbers))

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			results <- num * num
		}
		close(results)
	}()

	return results
}

func main() {
	// Массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	wg := new(sync.WaitGroup)

	results := squares(wg, numbers)

	// Чтение результатов из канала и вывод в stdout
	for result := range results {
		fmt.Println(result)
	}
}
