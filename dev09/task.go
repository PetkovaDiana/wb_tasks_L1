package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)    // канал для записи
	chOut := make(chan int) // канал для вывода

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	square(&wg, ch, chOut)
	read(&wg, chOut)

	for _, val := range arr {
		ch <- val
	}
	close(ch)

	wg.Wait()

}

func square(wg *sync.WaitGroup, ch <-chan int, chOut chan<- int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			res := val * val
			chOut <- res
		}
		close(chOut)
	}()

}

func read(wg *sync.WaitGroup, chOut <-chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range chOut {
			fmt.Println(val)
		}
	}()
}
