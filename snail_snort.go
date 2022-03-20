package main

import (
	"fmt"
	"math"
)

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

	if len(arr[0]) == 1 {
		return arr[0]
	}

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

	max := int(math.Floor(float64(len(arr) / 2)))
	start := index

	if start > max {
		start = max
	}

	currArr := arr[start:len(arr) - index]

	if index > 0 {
		var nestArr [][]int
		for _, v := range currArr {
			nestArr = append(nestArr, v[start:len(v) - index])
		}
		currArr = nestArr
	}

	if len(acc) < size {
		curr := sort(currArr)
		return snail(arr, index + 1, append(acc, curr...), size)
	}

	return acc
}

func SnailSort(arr [][]int) []int {
	size := len(arr) * len(arr[0])
	return snail(arr, 0, []int{}, size)
}

func main() {

	fmt.Println("0 x 0 =>", SnailSort([][]int{{}}))

	arr0 := [][]int{
		{1},
	}

	fmt.Println("1 x 1 =>", SnailSort(arr0))

	arr1 := [][]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println("2 x 2 =>", SnailSort(arr1))

	arr2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("3 x 3 =>", SnailSort(arr2))

	arr3 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	fmt.Println("4 x 4 =>", SnailSort(arr3))

	arr4 := [][]int{
		{1,   2,  3,  4,  5,  6},
		{7,   8,  9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36},
	}

	fmt.Println("6 x 6 =>", SnailSort(arr4))

	arr5 := [][]int{
		{ 1,  2,  3,  4, 5},
		{ 6,  7,  8,  9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	fmt.Println("5 x 5 =>", SnailSort(arr5))
}
