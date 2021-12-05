package main

import (
	"fmt"
	"strings"
)

func Reverse(s []string) string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, " ")
}

func main() {
	str := "Hello mans"
	mass := strings.Split(str, " ")
	fmt.Println(Reverse(mass))
}
