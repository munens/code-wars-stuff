package main

import "fmt"

/* 1st attempt:
import (
	"fmt"
	"math")

func getThresholds(size int) (thresholds []int) {
	sqrt := math.Sqrt(float64(size))
	for i := 1; i <= size; i++ {
		if i % int(sqrt) == 0 {
			thresholds = append(thresholds, i)
		}
	}

	return thresholds
}

func getIndexes(arr [][]int) (xIndexes []int) {
	colLen := len(arr)
	count :=  colLen * colLen
	thresholds := getThresholds(count)
	maxIndex := colLen - 1

	fmt.Println(thresholds)

	for i := 0; i < len(thresholds); i++ {
		for j := 0; j < count; j++ {

			if j < thresholds[0] {
				xIndexes = append(xIndexes, 0)
			}

			if j < thresholds[1] {
				if j < maxIndex {
					//xIndexes = append(xIndexes, j - maxIndex)
				}
			}
		}
	}

	return xIndexes
}

func snailSort(arr [][]int) {

	fmt.Println(getIndexes(arr))

}

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	snailSort(arr)
}


 */


/* 2nd attempt */

func filter(compare []int, current []int) {

}

func concat(curr []int, next[][]int) []int {

	if len(next) == 0 {
		return curr
	}

	nextCurr := append(curr, next[0]...)
	return concat(nextCurr, next[1:])
}

func reverse(arr []int) (rev []int ) {
	for i := len(arr) - 1; i >= 0; i-- {
		rev = append(rev, arr[i])
	}
	return rev
}

func sort(arr [][]int) []int {
	firstRow := arr[0]

	var middleRows [][]int
	for rowIdx := 1; rowIdx < len(arr) - 1; rowIdx++ {
		middleRows = append(middleRows, arr[rowIdx])
	}

	var rightRow []int
	for rowIdx := 0; rowIdx < len(middleRows); rowIdx++ {
		currRow := middleRows[rowIdx]
		rightRow = append(rightRow, currRow[len(currRow) - 1])
	}

	var leftRow []int
	for rowIdx := 0; rowIdx < len(middleRows); rowIdx++ {
		currRow := middleRows[rowIdx]
		leftRow = append(leftRow, currRow[0])
	}

	bottomRow := arr[len(arr) - 1]

	return concat(
		firstRow,
		[][]int{
			rightRow,
			reverse(bottomRow),
			reverse(leftRow),
		},
	)
}

func snail(arr [][]int, index int, acc []int, size int) []int {

	currArr := arr[index:len(arr) - index]

	if index > 0 {
		var nestArr [][]int
		for _, v := range currArr {
			nestArr = append(nestArr, v[index:len(v) - index])
		}
		currArr = nestArr
	}
	fmt.Println(acc)


	if len(acc) < size {
		curr := sort(currArr)
		return snail(arr, index + 1, append(acc, curr...), size)
	}

	return acc
}

func SnailSort(arr [][]int) {
	size := len(arr) * len(arr[0])
	snail(arr, 0, []int{}, size)
}

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	SnailSort(arr)

	arr2 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	SnailSort(arr2)
}
