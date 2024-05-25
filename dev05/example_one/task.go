package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

// Первый пример через context

func main() {
	//Создается контекст с тайм-аутом 1 секунда. Контекст автоматически отменяется по истечении времени.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	writeChan := write(ctx, &wg)
	read(writeChan, &wg)

	wg.Wait()
}

func write(ctx context.Context, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		var k int
		defer wg.Done()
	loop:
		for {
			select {
			//Если контекст ctx завершен (ctx.Done()), канал ch закрывается и выполнение цикла прерывается (break loop)
			case <-ctx.Done():
				close(ch)
				break loop
				// return
			default:
				//Если контекст не завершен, горутина засыпает на 200 миллисекунд и затем отправляет значение k в канал ch, после чего увеличивает k.
				time.Sleep(200 * time.Millisecond)
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

		for val := range writeChan {
			fmt.Println(val)
		}
	}()
}
