package main

import (
	"fmt"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	defer close(ch)
	for i := 0; i < 5; i++ {
		go func() {
			ch <- arr[i] * arr[i]
		}()
		fmt.Println(<-ch)
	}
}
