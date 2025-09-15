package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var n int // Создали переменную для сканера

	fmt.Println("Введите N воркеров:")
	fmt.Scan(&n) //Считали данные с консоли, записали в n

	ch := make(chan int)

	//Создаем канал для прерываний
	osSignal := make(chan os.Signal, 1) // Сигнал SIGINT отправляется при введении пользователем в управляющем терминале символа прерывания, по умолчанию это ^C (Control-C).
	// перехватыват сигнала. signal.Notify регистрирует данный канал для получений уведомлений об определенных сигналах от ОС
	signal.Notify(osSignal, syscall.SIGINT)

	read(ch, n)

	write(ch, osSignal)
}

func read(ch chan int, n int) {
	for i := 0; i < n; i++ {
		go func() {
			for num := range ch { // Читает из канала, пока он не закрыт
				fmt.Println(num)
			}
		}()
	}
}

func write(ch chan int, osSignal chan os.Signal) {
	for {
		select { //С помощью неблокирующего оператора select читаем канал
		case <-osSignal:
			close(ch)
			fmt.Println("Завершено")
			return // Выходим из функции
		default:
			ch <- 2
		}
	}
}
