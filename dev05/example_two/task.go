package main

import (
	"fmt"
	"sync"
	"time"
)

// Это второй способ с использование канала done,
//который по истечению определенного времени закрывается

func main() {
	var wg sync.WaitGroup
	done := closer()

	writeChan := write(done, &wg)
	read(writeChan, &wg)

	wg.Wait()
}

func write(done <-chan struct{}, wg *sync.WaitGroup) <-chan int {
	// Создае канал для передачи типов int
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		var k int
		for {
			select {
			case <-done:
				//Если канал done закрыт, закрывается канал ch и горутина завершает работу.
				close(ch)
				return
			default:
				// Если канал done не закрыт, горутина засыпает на 500 миллисекунд и затем отправляет значение k в канал ch,
				// после чего увеличивает k
				time.Sleep(500 * time.Millisecond)
				ch <- k
				k++
			}
		}
	}()

	return ch
}

func read(writeChan <-chan int, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		//В бесконечном цикле for читаются значения из writeChan и выводятся на экран.
		// Цикл завершится, когда writeChan будет закрыт
		for val := range writeChan {
			fmt.Println(val)
		}
	}()
}

func closer() <-chan struct{} {
	done := make(chan struct{})

	//Запускается горутина, которая засыпает на 2 секунды, после чего закрывает канал done
	go func() {
		time.Sleep(2000 * time.Millisecond)
		close(done)
	}()

	return done
}
