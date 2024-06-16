package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		a, b, p big.Int
		op      string
	)

	fmt.Scan(&a, &b, &op)

	switch op {
	case "+":
		fmt.Println(p.Add(&a, &b))
	case "-":
		fmt.Println(p.Sub(&a, &b))
	case "*":
		fmt.Println(p.Mul(&a, &b))
	case "/":
		fmt.Println(p.Div(&a, &b))
	}
}
