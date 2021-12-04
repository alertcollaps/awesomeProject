package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	n := 7
	mp := new(mp)
	mp.mp = make(map[int]int)
	f := func(i int) {
		mp.Lock()
		defer mp.Unlock()
		mp.mp[rand.Intn(10)] = rand.Intn(10)
		fmt.Printf("GoRutine №%v, closed\n", i)
	}
	fmt.Println("Please wait")
	fmt.Println("For continue press Enter")
	for i := 0; i < n; i++ {
		go f(i)
	}
	fmt.Scanln()
	fmt.Println("Map generated:", mp.mp)
}

type mp struct {
	sync.Mutex
	mp map[int]int
}
