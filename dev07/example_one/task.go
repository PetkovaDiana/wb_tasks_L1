package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

// Первый пример как можно решить задачу

var (
	mx sync.Mutex     // Создание мьютекса для синхронизации доступа к разделяемым данным
	wg sync.WaitGroup // Создание группы ожидания для ожидания завершения всех горутин
)

func main() {

	mp := make(map[int]int) // Создание пустого среза для хранения данных

	wg.Add(10) // Добавление 10 ожидаемых завершений в группу ожидания для 10 горутин
	for i := 0; i < 10; i++ {
		i := i // Создание копии переменной i для каждой итерации цикла
		go func() {
			defer wg.Done()   // Уменьшение счетчика группы ожидания при завершении горутины (3)
			mx.Lock()         // Блокировка мьютекса перед изменением среза (1)
			defer mx.Unlock() // Разблокировка мьютекса после завершения операции (4)
			mp[i] = i         // Добавление элемента в срез (2)
		}()
	}

	wg.Wait()       // Ожидание завершения всех горутин
	fmt.Println(mp) // Вывод содержимого среза arr после завершения всех операций
}
