package main

import (
	"fmt"
	"math"
)

func main() {
	constant := math.MaxInt64
	var x int64
	fmt.Scanln(&x)
	fmt.Printf("%064b(%d)\n", x, x)
	var bit int
	fmt.Print("Enter number of bit: ")
	fmt.Scanln(&bit)
	tmp := 1 << bit
	x = x ^ int64(tmp)
	if x < 0 { // tracking overflow
		x += 1
		x = (x ^ int64(constant>>1)) + 2
	}
	fmt.Printf("%064b(%d)\n", x, x)

}
