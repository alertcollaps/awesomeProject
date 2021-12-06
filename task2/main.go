package main

import (
	"fmt"
	"sync"
)

func write(c chan int, val int, wt *sync.WaitGroup) {
	c <- val
	defer wt.Done()
}

func main() {
	var c = make(chan int)
	var mass = []int{2, 4, 6, 8, 10}
	wt := sync.WaitGroup{}
	wt.Add(len(mass))
	go func() {
		for _, val := range mass {
			go write(c, val, &wt)
		}
		wt.Wait()
		close(c)
	}()
	for {
		val, open := <-c
		if !open {
			break
		}
		fmt.Println(val * val)
	}

}
