package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	Exit := make(chan int)
	var n int = 1
	go func() {
		for {
			select {
			case c <- rand.Intn(10):
			case <-Exit:
				fmt.Println("Iteration closed")
				Exit <- 0
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case mess := <-c:
				fmt.Println("Read message:", mess)
			case <-Exit:
				fmt.Println("Read closed")
				Exit <- 0
				return
			}
		}
	}()
	<-time.After(time.Second * time.Duration(n))
	Exit <- 0
	time.Sleep(time.Millisecond)
}
