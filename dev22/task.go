package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация больших чисел a и b
	a := new(big.Int)
	b := new(big.Int)

	// Присваивание значений a и b (например, числа больше 2^20)
	a.SetString("1048577", 10) // 2^20 + 1
	b.SetString("2097153", 10) // 2^21 + 1

	// Выполнение арифметических операций
	sum := new(big.Int).Add(a, b)
	difference := new(big.Int).Sub(a, b)
	product := new(big.Int).Mul(a, b)
	quotient := new(big.Int).Div(a, b)

	// Вывод результатов
	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Printf("a + b = %s\n", sum.String())
	fmt.Printf("a - b = %s\n", difference.String())
	fmt.Printf("a * b = %s\n", product.String())
	fmt.Printf("a / b = %s\n", quotient.String())
}
