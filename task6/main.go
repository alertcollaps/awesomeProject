package main

import (
	"fmt"
	"sync"
	"time"
)

func gorut1(n int) {
	for {
		n--
		fmt.Println("Gorutine1 write", n)
		if n == 0 {
			fmt.Println("Gorutine1 stoped")
			return
		}
	}
}
func gorut2(Exit chan bool, read chan int) {
	for {
		select {
		case m := <-read:
			fmt.Println("Gorutine2 Read", m)

		case <-Exit:
			fmt.Println("Gorutine2 was closed")
			return
		}
	}
}
func gorut3(wt *sync.WaitGroup, read chan int) {
	for {
		if data, boll := <-read; boll {
			fmt.Println("Gorutine3 read", data)
		} else {
			fmt.Println("Gorutine3 closed")
			wt.Done()
			return
		}
	}
}
func main() {
	Exit := make(chan bool)
	Data := make(chan int)
	wt := sync.WaitGroup{}
	wt.Add(1)
	go gorut1(5)
	go gorut2(Exit, Data)
	go gorut3(&wt, Data)
	go func() {
		for i := 0; i < 15; i++ {
			Data <- 1
			Data <- 2
			Data <- 3
			time.Sleep(time.Millisecond)
		}
	}()
	time.Sleep(1 * time.Second)
	Exit <- false
	time.Sleep(1 * time.Second)
	close(Data)
	wt.Wait()
	fmt.Println("Main closed")
}
