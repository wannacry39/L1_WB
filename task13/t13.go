package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a, b int64
	fmt.Scan(&a, &b)

	fmt.Println(a, b)
	// a, b = b, a
	a = atomic.SwapInt64(&a, b) //atomic realization
	fmt.Println(a, b)

}
