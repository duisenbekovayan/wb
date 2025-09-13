# Завершение программы по Ctrl+C (SIGINT)

Этот пример демонстрирует, как корректно завершать работу Go-программы и всех горутин-воркеров при нажатии **Ctrl+C**.

## 📖 Описание

- Программа запускает несколько воркеров в отдельных горутинах.
- Каждый воркер выполняет работу в цикле и проверяет, не пришёл ли сигнал завершения.
- При нажатии **Ctrl+C** (SIGINT) программа:
    1. Ловит сигнал от ОС.
    2. Отправляет отмену через `context.Context`.
    3. Дожидается завершения всех горутин с помощью `sync.WaitGroup`.
    4. Корректно выходит.

Использование `context.Context` – идиоматичный способ в Go, который позволяет централизованно управлять временем жизни горутин и предотвращать утечки.

## 📝 Код примера

```go
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
	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	<-sigChan
	fmt.Println("\nПолучен Ctrl+C, завершаем работу...")
	cancel()
	wg.Wait()

	fmt.Println("Все воркеры корректно завершились.")
}
