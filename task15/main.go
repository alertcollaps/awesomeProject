package main

import "fmt"

var justString string

//Может случиться такая ошибка "panic: runtime error: slice bounds out of range [:100] with length {<100}"
func someFunc() {
	var l int = 100
	v := createHugeString(1 << 10)
	//Добавляем проверку на размер слайса
	if l > len(v) {
		l = len(v)
	}
	justString = v[:l]
}

func main() {
	someFunc()
}
