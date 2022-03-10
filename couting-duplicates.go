package main

import (
	"fmt"
	"strings"
)

func duplicateCount(str string) int {
	arr := strings.Split(strings.ToLower(str), "")
	dict := make(map[string]int)
	for i := 0; i < len(arr);  i++ {
		char := arr[i]

		if ok := dict[char]; ok > 0 {
			currVal := dict[char] + 1
			dict[char] = currVal
		}

		if ok := dict[char]; ok == 0 {
			dict[char] = 1
		}
	}

	distinct := 0

	for _, v := range dict {
		if v > 1 {
			distinct += 1
		}
	}

	return distinct
}

func main() {
	fmt.Println(duplicateCount("aabbcsssSA122222"))
	fmt.Println(duplicateCount("abcde"))
	fmt.Println(duplicateCount("abcdea"))
	fmt.Println(duplicateCount("abcdeaB11"))
	fmt.Println(duplicateCount("indivisibility"))
}