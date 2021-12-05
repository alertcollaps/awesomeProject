package main

import (
	"fmt"
)

func find(a []int, val int) int {
	var mid int = len(a) / 2
	var left int = 0
	var right int = len(a) - 1
	for right-left != 0 {
		if val < a[mid] {
			right = mid
		} else if val > a[mid] {
			left = mid + 1
		} else {
			return mid
		}
		mid = left + (right-left)/2
	}
	return mid
}

func main() {
	slice := []int{1, 3, 9, 19, 23, 24, 25, 26, 89}
	fmt.Println(find(slice, 1))
}
