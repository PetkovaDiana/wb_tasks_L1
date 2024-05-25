package main

import (
	"log"
	"time"
)

// Второй способ с помощью log.Fatal

func worker() {
	go func() {
		for {
			log.Println("Worker is running...")
			time.Sleep(1 * time.Second)
		}
	}()
}

func main() {
	worker()

	// Ожидание 5 секунд перед вызовом log.Fatal
	time.Sleep(5 * time.Second)

	log.Fatal("Terminating program")
}
