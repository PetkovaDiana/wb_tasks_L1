package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

// В этом примере делаем все тоже самое как и в task02, только чуть по другому
// Функция для вычисления квадрата числа и отправки результата в канал,
// канал только для записи
func square(num int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	results <- num * num
}

func main() {
	// Массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Канал для передачи результатов
	results := make(chan int, len(numbers))

	// Группа ожидания для синхронизации горутин
	var wg sync.WaitGroup

	// Запуск горутин для вычисления квадратов чисел
	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, results)
	}

	// Закрытие канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Суммирование результатов из канала
	sum := 0
	for result := range results {
		sum += result
	}

	// Вывод суммы квадратов чисел
	fmt.Println("Sum of squares:", sum)
}
