package main

import "fmt"

func main() {
	var x int64
	fmt.Scanln(&x)
	fmt.Printf("%064b(%d)\n", x, x)
	var bit int
	fmt.Print("Enter number of bit: ")
	fmt.Scanln(&bit)
	tmp := 1 << bit
	x = x ^ int64(tmp)
	fmt.Printf("%064b(%d)\n", x, x)

}
