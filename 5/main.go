package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var n int
	ch := make(chan int, 1)
	fmt.Println("Wtire number:")
	fmt.Fscan(os.Stdin, &n) //Считали данные с консоли, записали в n
	go write(ch)
	go read(ch)

	<-time.After(time.Duration(n) * time.Microsecond) // Ждет пока истечет время, затем отправляет текущее время в канал
	fmt.Println("over")

}

// Функция бесконечно пишет в канал
func write(ch chan int) {
	for {
		ch <- 5
		fmt.Println("Write in channal")
	}
}

// Функция бесконечно читает из канала
func read(ch chan int) {
	for {
		<-ch
		fmt.Println("Reag from channal")
	}
}
