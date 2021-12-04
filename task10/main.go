package main

import (
	"fmt"
	"math"
)

func main() {
	mp := make(map[int][]float32)
	mass := [...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	min := mass[0]
	max := mass[0]
	for _, val := range mass {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	h := math.Abs(float64(max-min)) / 10
	for _, val := range mass {
		for j := 0; j < int(math.Ceil(h)); j++ {
			if val < (min+10+float32(10*j)) && val >= (min+float32(10*j)) {
				mp[j] = append(mp[j], val)
				break
			}
		}
	}
	fmt.Println(mp)
}
