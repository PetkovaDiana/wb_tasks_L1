package main

import "fmt"

func set(s []string) []string {
	var res []string
	setOfString := make(map[string]string, len(s))

	for _, val := range s {
		setOfString[val] = val
	}

	for key := range setOfString {
		res = append(res, key)
	}

	return res
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	setArr := set(arr)

	fmt.Println(setArr)
}
