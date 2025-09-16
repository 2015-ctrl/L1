package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	num int
	sync.Mutex
}

func (c *Counter) Inc() { //метод структуры счетчика, прибавляющая значение к счетчику
	c.Lock()         // как только одна из горутин вызывает метод Lock, другие не могут работать над структурой
	defer c.Unlock() // разблокировка после выполнения действия
	c.num += 2
}

func (c *Counter) Value() int { // метод возвращает значение счетчика
	return c.num
}

func main() {

	wg := new(sync.WaitGroup)

	c := &Counter{ // создаем объект структуры
		num: 0,
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()

	fmt.Println(c.Value())
}
