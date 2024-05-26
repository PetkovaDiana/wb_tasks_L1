package main

import (
	"fmt"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

*/

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := TmGroup(temperatures)

	fmt.Println("Группы температур:", groups)
}

func TmGroup(tmp []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, val := range tmp {
		key := int(val/10) * 10
		groups[key] = append(groups[key], val)
	}
	return groups
}
