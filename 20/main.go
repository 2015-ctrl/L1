package main

import (
	"fmt"
)

func reverseRunes(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

func reverseWords(s string) string {
	runes := []rune(s)

	//Переворачиваем всю строку
	reverseRunes(runes, 0, len(runes)-1)

	//Переворачиваем каждое слово отдельно
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' { // нашли конец слова
			reverseRunes(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	s := "snow dog sun"
	fmt.Println("Вход:  ", s)
	fmt.Println("Выход: ", reverseWords(s))
}
