package main

import "fmt"

func quicksort(a []int, left int, right int) []int {
	if left < right {
		pivot := a[left]

		i, j := left, right // Заводим два счётчика: left идёт слева, right идёт справа
		for {
			for a[j] >= pivot && i < j { // Уменьшаем j пока элементы больше пивота
				//fmt.Println("rigth", a[j])
				j--
			}

			for a[i] <= pivot && i < j { // Увеличиваем i пока элементы меньше пивота
				//fmt.Println("left", a[i])
				i++
			}

			if i >= j { // когда счетчики встретятся выходим из цикла
				break
			}

			a[i], a[j] = a[j], a[i] // меняем местами элементы
		}

		a[left] = a[i]
		a[i] = pivot

		quicksort(a, left, i-1) // Рекурсивно сортируется отдельно слева и правые стороны
		quicksort(a, i+1, right)
	}
	return a
}

func main() {
	a := []int{1, 22, 32, 62, 100, 10, 42, 52, 1} // 8

	fmt.Println(quicksort(a, 0, len(a)-1))
}
