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
	b    bool
}

func worker(c chan *Data, n int) {
	for {
		t, open := <-c
		if !open {
			fmt.Printf("Woker%d  closed\n", n)
			break
		}
		fmt.Printf("Woker%d  %d\n", n, t.data)
	}

}
func close1(c chan *Data) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("exiting")
	t := <-c
	t.b = false
	close(c)

}
func main() {
	c := make(chan *Data)
	dat := Data{1, true}

	var n int = 8
	go close1(c)
	for i := 0; i < n; i++ {
		go worker(c, i)
	}
	for {
		time.Sleep(time.Second)
		if !dat.b {
			break
		}
		dat.data = rand.Intn(10)
		c <- &dat
	}

}
