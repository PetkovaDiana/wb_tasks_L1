package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
*/
func checkUnique(str string) bool {

	str = strings.ToLower(str)

	charMap := make(map[rune]bool)

	for _, char := range str {

		if charMap[char] {
			return false
		}

		charMap[char] = true
	}

	return true
}

func main() {

	testStrings := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, str := range testStrings {
		fmt.Printf("All characters in \"%s\" are unique: %t\n", str, checkUnique(str))
	}
}
