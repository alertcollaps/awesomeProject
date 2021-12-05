package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	sync.Mutex
}

func main() {
	ct := counter{}
	n := 9
	wt := sync.WaitGroup{}
	wt.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			ct.count++
			wt.Done()
		}()
	}
	wt.Wait()
	fmt.Println("Значение счетчика равно:", ct.count)
}
