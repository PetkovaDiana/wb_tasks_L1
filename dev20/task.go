package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {

	srt := "snow dog sun"

	fmt.Println(srt)

	outStr := reverseWords(srt)

	fmt.Println(outStr)

}

func reverseWords(str string) string {
	//разбиваем строку на подстроки по пробельным символам и возвращает срез этих подстрок
	words := strings.Fields(str)

	left, right := 0, len(words)-1

	if left <= right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	// объединяем их
	return strings.Join(words, " ")
}
