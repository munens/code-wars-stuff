//Your task is to split the chocolate bar of given dimension n x m into small squares. Each square is of size 1x1 and unbreakable. Implement a function that
// will return minimum number of breaks needed.
//
//For example if you are given a chocolate bar of size 2 x 1 you can split it to single squares in just one break, but for size 3 x 1 you must do two breaks.
//
//If input data is invalid you should return 0 (as in no breaks are needed if we do not have any chocolate to split). Input will always be a non-negative integer.

package main

import "fmt"

func break_chocolate(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	return m * n - 1
}

func main() {
	fmt.Println(fmt.Sprintf("%d x %d = %d", 1, 1, break_chocolate(1, 1)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 2, 2, break_chocolate(2, 2)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 3, 3, break_chocolate(3, 3)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 4, 4, break_chocolate(4, 4)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 5, 5, break_chocolate(5, 5)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 6, 6, break_chocolate(6, 6)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 7, 7, break_chocolate(7, 7)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 7, 1, break_chocolate(7, 1)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 6, 1, break_chocolate(6, 1)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 5, 1, break_chocolate(5, 1)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 4, 1, break_chocolate(4, 1)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 3, 1, break_chocolate(3, 1)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 2, 1, break_chocolate(2, 1)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 7, 6, break_chocolate(7, 6)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 7, 5, break_chocolate(7, 5)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 7, 4, break_chocolate(7, 4)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 7, 3, break_chocolate(7, 3)))
	fmt.Println(fmt.Sprintf("%d x %d = %d", 7, 2, break_chocolate(7, 2)))
}