package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Data struct {
	data int
}

func worker(c chan *Data, n int, Exit chan int) {
	for {
		select {
		case data := <-c:
			fmt.Printf("Woker%d  %d\n", n, data.data)
		case <-Exit:
			fmt.Printf("Woker%d  closed\n", n)
			return
		}
	}

}
func close1(Exit chan int, n int) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("exiting")
	for i := 0; i < n+1; i++ {
		Exit <- 0
	}
}

func main() {
	c := make(chan *Data)
	dat := Data{1}
	Exit := make(chan int)
	var n int = 8
	go close1(Exit, n)
	for i := 0; i < n; i++ {
		go worker(c, i, Exit)
	}
	for {
		time.Sleep(time.Second)
		dat.data = rand.Intn(10)
		select {
		case c <- &dat:
		case <-Exit:
			time.Sleep(time.Millisecond)
			fmt.Printf("Main closed\n")
			return
		}

	}

}
