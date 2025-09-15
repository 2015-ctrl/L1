package main

import (
	"sync"
)

var wg sync.WaitGroup //для ожидания завершения группы горутин

func main() {
	var user sync.Map //потокобезопасная структура данных

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			user.Store("vasa", 11) //Метод Store используется для безопасной записи пары ключ-значение
		}()
	}
	wg.Wait()
}
