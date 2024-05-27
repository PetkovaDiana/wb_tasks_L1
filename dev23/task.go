package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}

	target := 5 // индекс, который надо нам удалить

	arrNew := deleteArr(arr, target)

	fmt.Println(arrNew)
}

func deleteArr(arr []int, target int) []int {

	// удаляем через доабвление среза
	arr = append(arr[:target], arr[target+1:]...)

	return arr
}
