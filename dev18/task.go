package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

// Counter структура, которая содержит счетчик и мьютекс для защиты доступа к нему
type Counter struct {
	sync.Mutex
	counter int
}

// Создаем экземпляр структуру Counter
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.Lock()
	c.counter++
	c.Unlock()
}

func (c *Counter) String() string {
	return fmt.Sprint(c.counter)
}

func main() {

	c := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c)
}
