package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

// проблемы этого кода: утечка памяти, неправильное использование глобальной переменной
// createHugeString создает большую строку из случайных символов.
func createHugeString(size int) string {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < size; i++ {
		sb.WriteRune('a' + rune(generator.Intn('z'-'a'+1)))
	}
	return sb.String()
}

// someFunc обрабатывает строку и возвращает первые 100 символов.
func someFunc() string {
	v := createHugeString(1 << 10)

	// Преобразуем строку в срез рун
	runes := []rune(v)

	// Проверяем, что длина строки больше 100 символов
	if len(runes) > 100 {
		runes = runes[:100]
	}

	// Создаем новую строку из рун, чтобы избежать утечки памяти
	return string(runes)
}

func main() {
	justString := someFunc()
	fmt.Println("justString =", justString)
}
