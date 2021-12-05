package main

import (
	"fmt"
	"strings"
)

func unic(s string) bool {
	s = strings.ToLower(s)
	testMap := make(map[rune]int)
	for _, val := range s {
		testMap[val]++
		if testMap[val] == 2 {
			return false
		}
	}
	return true
}

func main() {
	str := "HeИloПривет"
	fmt.Println(unic(str))
}
