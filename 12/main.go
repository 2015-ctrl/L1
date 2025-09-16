package main

import "fmt"

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	b := make(map[string]bool) // Мапа, где ключи — это строки из массива, а значения — булевы флаги
	for _, v := range a {      //Итерация по элементам массива и запись их в мапу как ключей с присвоением каждому ключу значения true
		b[v] = true
	}

	// Выводим уникальные строки
	for k := range b {
		fmt.Println(k)
	}
}
