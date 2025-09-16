package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// функция для умножения чисел на два
func square(in chan int, out chan int) {
	defer wg.Done()
	for val := range in { //интерация по каналу in, на каждой итерации текущий поток блокируется до следующего значения в канале
		square := val * 2
		out <- square
	}

	close(out) // новых значений в канал поступать не будет

}

func read(out chan int) { //считываение данных из канала out
	defer wg.Done()
	for val := range out { //интерация по каналу out и вывод его переменных
		fmt.Println(val)
	}

}

func main() {

	in := make(chan int, 3)
	out := make(chan int, 3)

	a := [3]int{1, 2, 3}

	wg.Add(3)

	go func(in chan int) {
		defer wg.Done()
		for _, val := range a { // цикл который передает все значения в канал in
			in <- val
		}
		close(in)
	}(in)

	go square(in, out) // горутина которая умножает числа на два канала in и передает их в канал out

	go read(out) // горунта которая выводит значения из канала out

	wg.Wait()

}
