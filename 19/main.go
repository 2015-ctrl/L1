package main

import "fmt"

func main() {
	s := "главрыба"
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	r := []rune(s) // Создаем слайс рун по строке
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
