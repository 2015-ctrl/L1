package main

import (
	"fmt"
	"time"
)

// 1)Возврат из функции: Когда функция, запущенная в goroutine, завершает свое выполнение, goroutine завершается. Таким образом, можно просто выйти из функции, чтобы завершить goroutine.
func main() {
	go goroutine()
	time.Sleep(time.Millisecond * 1300)
	fmt.Println("Finish")

}

func goroutine() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 5)
		fmt.Println(i)
	}
	return
}
