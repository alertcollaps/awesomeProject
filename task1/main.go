package main

import (
	"fmt"
)

type Human struct {
	height int
	weight int
}

type Action struct {
	sit   bool
	sleep bool
	Human //Наследование
}

//Метод структуры Human
func (h Human) index() float32 {
	return float32((h.height - 100) / h.weight)
}

func main() {
	//Использование собственных методов и полей структуры Human
	human := new(Human)
	human.height = 175
	human.weight = 75
	fmt.Printf("Struct(Human)\tindex = %.3f\n", human.index())
	//Использование как собственных полей, так и наследованных полей и метода index() структуры Action
	actHuman := new(Action)
	actHuman.height = 175
	actHuman.weight = 75
	actHuman.sit = true
	actHuman.sleep = false
	fmt.Printf("Struct(Action)\tindex = %.3f\n", actHuman.index()) //index() - наследованный метод
}
