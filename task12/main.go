package main

import "fmt"

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	mp := make(map[string]int)
	slice := []string{}
	for _, val := range str {
		if mp[val] == 1 {
			continue
		}
		mp[val]++
		slice = append(slice, val)
	}
	fmt.Println(slice)
}
