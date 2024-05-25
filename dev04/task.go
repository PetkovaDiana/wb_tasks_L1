package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Парсинг аргументов командной строки
	// По умолчанию будет использовано 5 воркеров, чтобы изменить,
	// нужно будет на старте указать флаг и кол-во воркеров (-workers 10)
	numWorkers := flag.Int("workers", 5, "number of workers")
	flag.Parse()

	// Канал для передачи данных
	jobs := make(chan int)

	// Канал для сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Группа ожидания для синхронизации горутин воркеров
	var wg sync.WaitGroup

	// Запуск воркеров
	// В этом цикле мы берем указатель на numWorkers, потому что numWorkers - это указатель на значение,
	// которое хранит кол-во воркеров, заданное через аргументы командной строки.
	// При использовании flag.Int мы создаем флаг workers, который будет хранить значение количества воркеров.
	// Функция flag.Int возвращает указатель на это значение
	wg.Add(*numWorkers)
	for i := 1; i <= *numWorkers; i++ {
		go worker(i, jobs, &wg)
	}

	// Генерация данных в главном потоке
	go processWorkers(stop, jobs)

	// Ожидание завершения работы всех воркеров
	wg.Wait()
	fmt.Fprintln(os.Stdout, "All workers have finished their work. Exiting.")
}

// Функция воркера для чтения данных из канала и вывода их в stdout
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Fprintf(os.Stdout, "Worker %d received job: %d\n", id, job)
	}
}

// Функция процесс воркера
func processWorkers(stop <-chan os.Signal, jobs chan<- int) {
	jobID := 1
	for {
		select {
		case <-stop:
			// Когда получим сигнал, то мы закроем канал jobs
			close(jobs)
			return
		default:
			// Отправка данных в канал
			jobs <- jobID
			jobID++
			time.Sleep(1 * time.Second) // имитация времени генерации данных
		}
	}

}
