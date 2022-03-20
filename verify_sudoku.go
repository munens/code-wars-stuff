package main

import "fmt"

func isUnique(arr []int) bool {
	dict := make(map[int]int)
	for _, v := range arr {
		if _, ok := dict[v]; !ok {
			dict[v] = 1
		} else {
			dict[v]++
		}
	}

	for _, v := range dict {
		if v > 1 {
			return false
		}
	}

	return true
}

func every(arr ...bool) bool {
	i := 0
	for i < len(arr) {
		if !arr[i] {
			return false
		}
	}

	return true
}

var size = 9

func verifyRows(arr [][]int) bool {
	for _, v := range arr {
		if ok := isUnique(v); !ok {
			return false
		}
	}

	return true
}

func verifyColumns(arr [][]int, size int) bool {
	var columns [][]int
	currX := 0
	for len(columns) < size {
		var currCol []int
		for y := 0; y < size; y++ {
			currCol = append(currCol, arr[y][currX])
		}
		columns = append(columns, currCol)
		currX++
	}

	for _, v := range columns {
		if ok := isUnique(v); !ok {
			return false
		}
	}

	return true
}

func take3x3s(arr [][]int) (all3x3s [][][]int) {
	thresh := 3
	xMin := 0
	yMin := 0
	yMax := thresh

	for len(all3x3s) < size {
		var curr3x3 [][]int
		var curr3x3Row []int
		currX := xMin
		for y := yMin; y < yMax; y++ {
			curr3x3Row = append(curr3x3Row, arr[currX][y])
		}
		curr3x3 = append(curr3x3, curr3x3Row)
		currX++

		if len(curr3x3) == 3 {
			yMin, yMax, xMin = yMin + thresh, yMax + thresh, 0
			all3x3s = append(all3x3s, curr3x3)
		}

		if len(all3x3s) % 3 == 0 {
			xMin = xMin + thresh
			yMin = 0
			yMax = thresh
		}
	}

	return all3x3s
}

func verify3x3(arr [][]int) bool {

	if ok := verifyColumns(arr, 3); !ok {
		return false
	}

	if ok := verifyRows(arr); !ok {
		return false
	}

	dict := make(map[int]int)
	currX := 0

	for len(dict) < size {
		yMin := 0
		yMax := 3
		for y := yMin; y < yMax; y++ {
			k := arr[currX][y]
			if _, ok := dict[k]; !ok {
				dict[k] = 1
			} else {
				dict[k]++
			}
		}

		currX++
	}

	for _, v := range dict {
		if v > 1 {
			return false
		}
	}

	return true
}

func verify3x3s(arr [][]int) bool {

	for _, v := range take3x3s(arr) {
		if ok := verify3x3(v); !ok {
			return false
		}
	}

	return true
}

func VerifySudoku(arr [][]int) bool {
	return every(verifyRows(arr), verifyColumns(arr, size), verify3x3s(arr))
}

func main() {
  puzzle := [][]int{
	  {5, 3, 4, 6, 7, 8, 9, 1, 2},
	  {6, 7, 2, 1, 9, 5, 3, 4, 8},
	  {1, 9, 8, 3, 4, 2, 5, 6, 7},
	  {8, 5, 9, 7, 6, 1, 4, 2, 3},
	  {4, 2, 6, 8, 5, 3, 7, 9, 1},
	  {7, 1, 3, 9, 2, 4, 8, 5, 6},
	  {9, 6, 1, 5, 3, 7, 2, 8, 4},
	  {2, 8, 7, 4, 1, 9, 6, 3, 5},
	  {3, 4, 5, 2, 8, 6, 1, 7, 9},
  }

  fmt.Println(puzzle, "verifySudoku => ", VerifySudoku(puzzle))
}
