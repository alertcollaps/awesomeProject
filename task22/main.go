package main

import (
	"fmt"
	"math"
	"math/big"
)

func operation(s string, a, b *big.Int) *big.Int {
	switch s {
	case "+":
		return a.Add(a, b)
	case "-":
		return a.Add(a, b.Neg(b))
	case "*":
		return a.Mul(a, b)
	case "/":
		return a.Div(a, b)
	default:
		return big.NewInt(0)
	}
}

func main() {
	a := big.NewInt(int64(math.Pow(2, 20)))
	b := big.NewInt(int64(math.Pow(2, 20)))
	fmt.Println(operation("*", a, b))
}
