package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func intersection(a, b []int) []int {
	setOne := make(map[int]int, len(a))
	setTwo := make(map[int]int, len(b))

	var res []int

	for _, val := range a {
		setOne[val] = val
	}

	for _, val := range b {
		setTwo[val] = val
	}

	for key := range setTwo {
		if setOne[key] == setTwo[key] {
			res = append(res, key)
		}
	}

	return res
}

func main() {
	setOne := []int{1, 3, 5, 7, 8, 9}
	setTwo := []int{2, 3, 4, 7, 5, 11}

	fmt.Println(intersection(setOne, setTwo))
}
