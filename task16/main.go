package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)

	quicksort(slice)

	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

	quicksort1(slice)

	fmt.Println("\n--- Sorted(2 func) ---\n\n", slice, "\n")
}

// Генерирует слайс размером size, элементы которого выбираются случайно в диапазоне от -999 до 999
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

//Фукнция сортировки, взята из одного сайта
func quicksort(a []int) []int {
	//Проверка, что длина слайса не меньше 2
	if len(a) < 2 {
		return a
	}
	//Опеределение границ слайса
	left, right := 0, len(a)-1
	//Определение индекса разделяющего элемента
	pivot := (left + right) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
func quicksort1(a []int) {
	//Проверка, что длина слайса не меньше 2
	if len(a) < 2 {
		return
	}
	var q = partition(a)
	quicksort1(a[:q])
	quicksort1(a[q+1:])

}

func partition(a []int) int {
	v := a[len(a)/2]
	i := 0
	j := len(a) - 1
	for i <= j {
		for a[i] < v {
			i++
		}
		for a[j] > v {
			j--
		}
		if i >= j {
			break
		}
		i++
		j--
		a[i], a[j] = a[j], a[i]
	}
	return j
}
