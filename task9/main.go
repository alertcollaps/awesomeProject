package main

import (
	"fmt"
	"sync"
)

func main() {
	//Объявление двух каналов конвееров
	get := make(chan int)
	post := make(chan int)
	//Объявление кол-ва элементов в массиве
	const n int = 8
	var mass [n]int
	//Заполение массива от 0 до n-1
	for i := 0; i < n; i++ {
		mass[i] = i
	}
	//Горутина, которая пишет в канал get
	go func() {
		for i := 0; i < len(mass); i++ {
			get <- mass[i]
		}
		//Закрытие канала после записи, а также условие окончание следующей горутины
		close(get)
		fmt.Println("Send to get ended")
	}()
	//Горутина, которая принимает данные из канала get и передает их в post
	go func() {
		for {
			if dat, ok := <-get; ok {
				post <- dat
			} else {
				fmt.Println("Send to post closed")
				close(post)
				return
			}
		}
	}()
	//Объявление waitgroup для последней горутины
	wt := sync.WaitGroup{}
	wt.Add(1)
	//Горутина, которая считывает данные из post и выводит их на экран
	go func() {
		for {
			if dat, ok := <-post; ok {
				fmt.Println("Get:", dat)
			} else {
				fmt.Println("Enter to screen ended")
				wt.Done()
				return
			}
		}
	}()
	//Ждем окончания последней горутины
	wt.Wait()
	fmt.Println("Main ended")
}
