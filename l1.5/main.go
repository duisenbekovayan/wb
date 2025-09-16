package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// Отправка чисел в канал
	go func() {
		i := 1
		for {
			ch <- i
			i++
			time.Sleep(500 * time.Millisecond) // имитация работы
		}
	}()

	// Таймаут N секунд
	timeout := time.After(5 * time.Second)

	for {
		select {
		case val := <-ch:
			fmt.Println("Получено:", val)
		case <-timeout:
			fmt.Println("Время вышло, завершаем работу.")
			return
		}
	}
}
