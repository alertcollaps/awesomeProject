package main

import "fmt"

func del(sl []int, i int) []int {
	sll := sl
	if len(sl) > i {
		return append(sll[:i], sll[i+1:]...)
	}
	return sl
}

func main() {
	slice := []int{1, 3, 4, 5, 7, 90}
	//Объявление i
	i := 2

	fmt.Println("Slice", del(slice, i))
}
