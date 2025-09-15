package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	ch := make(chan int, 5) //Создаем канал, куда будут отправляться квадраты чисел

	a := [5]int{2, 4, 6, 8, 10}

	wg.Add(1)
	go func(ch chan int) { //в анонимной горутине запускаем цикл, который будет записывать квадраты чисел
		defer wg.Done()
		for _, val := range a {
			ch <- val * val
		}

		close(ch) //закрытие канала, чтобы сигнализировать о завершении отправки данных
	}(ch)
	c := 0
	for nu := range ch { //Читаем канал
		c += nu //Прибавляем значение из канала в сумму

	}
	fmt.Println(c)

	wg.Wait()

}
