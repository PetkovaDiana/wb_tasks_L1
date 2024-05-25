package main

import (
	"context"
	"log"
	"time"
)

// Третий пример через context

func main() {
	// с помощью context.WithTimeout, мы можем установить время, через сколько сеунд закончится программа
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	worker(ctx)

	log.Fatal("Terminating program")
}

func worker(ctx context.Context) {
	log.Println("starting working...")

	for {
		select {
		// здесь мы проверяем case, если контекст завершен, то выводит в консоль
		case <-ctx.Done():
			log.Println("Worker is stopping...")
			return
		default:
			// если нет, то отрабатывается log.Println("Worker is running...")
			log.Println("Worker is running...")
			time.Sleep(1 * time.Second)
		}
	}

}
