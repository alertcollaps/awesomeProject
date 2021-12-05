package main

import "fmt"

func main() {
	//Есть два числа
	a := 3
	b := 5
	//При выполнении этих 4-ех операций, числа меняются местами
	a = a + b
	b = b - a
	b = -b
	a = a - b
	fmt.Printf("a = %v, b = %v", a, b)
}
