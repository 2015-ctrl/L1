package main

import (
	"fmt"
	"os"
)

func worker(id int, ch <-chan int) {
        // каждый воркер бесконечно читает данные из канала и выводит их
	for val := range ch {
		fmt.Printf("Worker %d got %d\n", id, val)
	}
}

func main() {
	var n int
	fmt.Print("Введите количество воркеров: ")
        // считываем количество воркеров, которые нужно запустить
	fmt.Fscan(os.Stdin, &n)

	ch := make(chan int)

	// запускаем N воркеров — каждый из них будет конкурентно читать данные из канала
	for i := 1; i <= n; i++ {
		go worker(i, ch)
	}

	// постоянная запись данных в канал (главная горутина)
	for i := 1; ; i++ {
		ch <- i
	}
}

