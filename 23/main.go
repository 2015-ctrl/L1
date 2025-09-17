package main

import "fmt"

func removeSliceOne(slice []int, index int) []int { //удаляет элемент из среза по индексу с использованием append
	return append(slice[:index], slice[index+1:]...)
}

func removeSliceTwo(slice []int, index int) []int { //удаляет элемент из среза по индексу с использованием функции copy
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func removeSliceTheree(slice []int, index int) []int { //удаляет элемент, используя цикл для сдвига элементов
	for i := index + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	return slice[:len(slice)-1]
}

func removeSliceFour(slice []int, index int) []int { //удаляет элемент, создавая новый срез и копируя в него данные
	result := make([]int, len(slice)-1)
	copy(result, slice[:index])
	copy(result[index:], slice[index+1:])
	return result
}

func main() {

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный срез:", slice)

	fmt.Println("Удаление с помощью append:", removeSliceOne(slice, 2))

	slice = []int{1, 2, 3, 4, 5}
	fmt.Println("Удаление с использованием copy:", removeSliceTwo(slice, 2))

	slice = []int{1, 2, 3, 4, 5}
	fmt.Println("Удаление с использованием цикла:", removeSliceTheree(slice, 2))

	slice = []int{1, 2, 3, 4, 5}
	fmt.Println("Удаление с использованием copy и созданием нового среза:", removeSliceFour(slice, 2))
}
