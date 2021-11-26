package main

import "fmt"

func write(c chan int, val int) {
	c <- val
}

func main() {
	var c chan int = make(chan int)
	sum := 0
	var mass = []int{2, 4, 6, 8, 10}
	go func() {
		for _, val := range mass {
			write(c, val)
		}
		close(c)
	}()
	for {
		val, open := <-c
		if !open {
			break
		}
		sum += val
	}
	fmt.Println(sum)
}
