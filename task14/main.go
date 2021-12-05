package main

import (
	"fmt"
	"reflect"
)

type tip interface {
}

//Первый способ, вручную проверять каждый тип
func Determinant(i tip) {
	if b, ok := i.(int); ok {
		fmt.Println("It's int. Value:", b)
	} else if b, ok := i.(string); ok {
		fmt.Println("It's string. Value:", b)
	} else if b, ok := i.(bool); ok {
		fmt.Println("It's bool. Value:", b)
	} else if b, ok := i.(chan int); ok {
		fmt.Println("It's chan int. Value:", b)
	}
}

func main() {
	var i tip = 9
	var j tip = make(chan float32)

	Determinant(i)
	//Второй способ, с помощью специальной функции
	fmt.Println(reflect.TypeOf(j))
}
