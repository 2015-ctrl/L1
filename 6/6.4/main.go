package main

import (
	"fmt"
	"runtime"
	"time"
)

// worker завершает работу с помощью runtime.Goexit()
func worker() {
	defer fmt.Println("defer внутри worker выполнится")
	fmt.Println("worker начал работу")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("worker вызывает runtime.Goexit()")
	runtime.Goexit() // завершает текущую горутину

	// Этот код никогда не выполнится
	fmt.Println("этого сообщения вы не увидите")
}

func main() {
	go worker()

	time.Sleep(1 * time.Second)
	fmt.Println("main завершен")
}
