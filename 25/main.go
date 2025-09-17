package main

import (
	"fmt"
	"time"
)

func sleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second) // Оправляет текущее время в канал по истечению установленного времени, затем функция читает из канала и завершается
}

func main() {
	fmt.Println("start")
	sleep(3)
	fmt.Println("end")
}
