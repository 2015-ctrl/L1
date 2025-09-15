package main

import (
	"context"
	"fmt"
	"time"
)

// 3)Использование context.Context

func myGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Ожидает закрытие канала Done(), который предоставляется объектом context.Context
			fmt.Println("Stopping goroutine")
			return
		default:
			// Выполнять работу в goroutine
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // инициализация контекста отмены ontext.WithCancel
	//  Эта функция cancel используется для отмены созданного контекста.
	defer cancel()

	go myGoroutine(ctx)

	// Ждем некоторое время
	time.Sleep(3000 * time.Millisecond)

	fmt.Println("Exiting main()")
}
