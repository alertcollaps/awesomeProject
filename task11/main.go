package main

import "fmt"

func main() {
	mp := make(map[int]int)
	mas1 := []int{1, 1, 8, 7, 3, 11, 123}
	mas2 := []int{87, 6, 8, 123, 1}

	for i := 0; i < len(mas1); i++ {
		if mp[mas1[i]] == 1 {
			continue
		}
		mp[mas1[i]]++
	}
	for i := 0; i < len(mas2); i++ {
		if mp[mas2[i]] == 2 {
			continue
		}
		mp[mas2[i]]++
	}
	slice := make([]int, 0, len(mas1))
	for i, val := range mp {
		if val == 2 {
			slice = append(slice, i)
		}
	}
	fmt.Println(slice)
}
