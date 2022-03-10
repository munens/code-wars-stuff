//Return the number (count) of vowels in the given string.
//
//We will consider a, e, i, o, u as vowels for this Kata (but not y).
//
//The input string will only consist of lower case letters and/or spaces.

package main

import (
	"fmt"
	"strings"
)

func vowelExists(v string) bool {
	arr := []string{"a", "e", "i", "o", "u"}
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func vowelCount(str string) int {
	count := 0
	arr := strings.Split(strings.Trim(str, " "), "")
	for _, v := range arr {
		if ok := vowelExists(v); ok {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(fmt.Sprintf("hello has %d vowels", vowelCount("hello")))
	fmt.Println(fmt.Sprintf("munene has %d vowels", vowelCount("munene")))
	fmt.Println(fmt.Sprintf("Shiwa has %d vowels", vowelCount("Shiwa")))
	fmt.Println(fmt.Sprintf("Kimathi and Nicole has %d vowels", vowelCount("Kimathi and Nicole")))
	fmt.Println(fmt.Sprintf("yolo has %d vowels", vowelCount("yolo")))
	fmt.Println(fmt.Sprintf("bad_ting has %d vowels", vowelCount("bad_ting")))

	// nice solution: https://www.codewars.com/kata/reviews/5bd0d703dfa73b037c000610/groups/5bd0d752abc66115eb002188
}