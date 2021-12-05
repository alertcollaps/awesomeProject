package main

import (
	"fmt"
	"time"
)

func sleep(t int) {
	tm := time.After(time.Duration(t) * time.Second)
	<-tm
}

func main() {
	fmt.Println("Start")
	sleep(5)
	fmt.Println("End")
}
