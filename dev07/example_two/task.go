package main

import (
	"fmt"
	"sync"
)

var (
	mp sync.Map
	wg sync.WaitGroup
)

func main() {
	wg.Add(10) // Добавление 10 ожидаемых завершений в группу ожидания для 10 горутин

	for i := 0; i < 10; i++ {
		i := i // Создание копии переменной i для каждой итерации цикла
		go func() {
			defer wg.Done() // Уменьшение счетчика группы ожидания при завершении горутины
			mp.Store(i, i)  // Добавление элемента в sync.Map
		}()
	}

	wg.Wait() // Ожидание завершения всех горутин

	// Вывод содержимого карты mp после завершения всех операций
	mp.Range(func(key, value interface{}) bool {
		fmt.Printf("%d: %d\n", key, value)
		return true
	})
}
