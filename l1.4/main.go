package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: завершаюсь...\n", id)
			return
		default:
			fmt.Printf("Worker %d: работаю...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Создаём контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Канал для сигналов ОС
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	// Запускаем несколько горутин
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// Ждём сигнала прерывания
	<-sigChan
	fmt.Println("\nПолучен Ctrl+C, завершаем работу...")

	// Отправляем сигнал отмены в контекст
	cancel()

	// Ждём завершения всех горутин
	wg.Wait()

	fmt.Println("Все воркеры корректно завершились.")
}
