package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// Первый способ с помощью сигнала ОС
// так мы можем остановить главный поток

func main() {
	// Создаем канал для сигналов с буфером размером 1
	//Буферизация канала размером 1 позволяет избежать блокировки отправителя сигнала, если приемник не успевает его обработать.
	c := make(chan os.Signal, 1)
	//Функция signal.Notify сообщает Go, что канал c должен получать уведомления о сигналах os.Interrupt и syscall.SIGTERM.
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Делаем бесконечный цикл
	for {
		select {
		case <-c: // если будет сигнал, то сработает case, цикл завершится и завершится программа
			fmt.Println("Bye")
			return
			//Каждую секунду выводится сообщение "Hello, world!", пока не будет получен сигнал.
		case <-time.After(1 * time.Second):
			fmt.Println("Hello, world!")
		}
	}
}
