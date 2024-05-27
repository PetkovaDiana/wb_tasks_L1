package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {

	fmt.Println("Start sleeping...")
	Sleep(2000) // Спим 2 секунды
	fmt.Println("Woke up!")
}

func Sleep(milliseconds int) {
	select {
	case <-time.After(time.Duration(milliseconds) * time.Millisecond):
	}
}
