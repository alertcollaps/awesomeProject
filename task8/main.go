package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//Ввод чисел
	var i int64
	fmt.Println("Введите число int64")
	fmt.Fscan(os.Stdin, &i)
	//z может быть от 0 до 63
	var z uint8
	fmt.Println("Введите бит")
	fmt.Fscan(os.Stdin, &z)
	//Проверка, чтобы бит было меньше 63, на всякий случай
	if z > 63 {
		fmt.Println("Бит больше 63")
		return
	}
	//Делаем рандомную функцию не повторяющейся
	rand.Seed(time.Now().UnixNano())
	//Объявления бита смещения
	var bits int64
	bits = 1
	//Смещение 1 на z
	bits = bits << z
	//Определение, есть ли в числе i на z-ом бите 0 или 1
	bits = i & bits
	//Генерация земены в 0 или 1 бит
	randm := rand.Intn(2)
	fmt.Printf("Замена %v-ого бита\n", randm)
	//Возвращение результата на то место, где нужно поменять бит
	bits = bits >> (z - uint8(randm))
	//Создание переменной чистки
	var clean int64
	clean = 1 << randm
	//Переменная clean очищает(делает нулевым) тот бит, в который нужно записать результат(bits)
	i = (i & ^clean) | bits
	fmt.Println(i)

}
