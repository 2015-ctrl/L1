package main

import (
	"fmt"
	"math"
)

// Функция группировки температур по диапазонам с шагом 10
func TGroup(tmp []float64) map[int][]float64 {
	groups := make(map[int][]float64) // карта: ключ = начало диапазона, значение = список температур

	for _, val := range tmp {
		var key int

		/*
			- Для положительных чисел используем Floor (округление вниз).
			  Пример: 19.0 → 19/10 = 1.9 → Floor(1.9) = 1 → 1*10 = 10.
			  Значит 19 попадает в диапазон 10..19.9.
			- Для отрицательных чисел используем Ceil (округление вверх).
			  Пример: -25.4 → -25.4/10 = -2.54 → Ceil(-2.54) = -2 → -2*10 = -20.
			  Значит -25.4 попадает в диапазон -20..-29.9.
		*/
		if val < 0 {
			key = int(math.Ceil(val/10) * 10)
		} else {
			key = int(math.Floor(val/10) * 10)
		}

		// Добавляем значение в соответствующую группу
		groups[key] = append(groups[key], val)
	}

	return groups
}

func main() {
	// Входные данные
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Группировка
	groups := TGroup(temperatures)

	// Вывод результата
	fmt.Println(groups)
}
