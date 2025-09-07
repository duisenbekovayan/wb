package main

import (
	"flag"
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d got job %d\n", id, job)
	}
}

func main() {
	// Параметр количества worker-ov
	n := flag.Int("n", 3, "number of workers")
	flag.Parse()

	jobs := make(chan int)

	// Запуск worker-ov
	for i := 1; i <= *n; i++ {
		go worker(i, jobs)
	}

	// пишем в канал
	jobID := 1
	for {
		jobs <- jobID
		jobID++
		time.Sleep(500 * time.Millisecond) // типа делаем работу
	}
}
