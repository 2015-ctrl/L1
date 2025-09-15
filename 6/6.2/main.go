package main

import (
	"fmt"
	"time"
)

//2)Использование канала

func myGoroutine(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Stopping goroutine")
			return
			// канал stopCh закрывается в основной горутине с помощью close(stopCh): когда читается из закрытого канала, выход из горутины выполняется с использованием return.
		default:
			// Выполнять работу в goroutine
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopCh := make(chan struct{})

	go myGoroutine(stopCh)

	// Ждем некоторое время
	time.Sleep(3000 * time.Millisecond)

	// Отправляем сигнал на завершение
	close(stopCh)

	fmt.Println("Exiting main()")
}
