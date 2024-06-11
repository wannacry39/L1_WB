package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	var res int
	for i := 0; i < 5; i++ {
		go func() {
			ch <- arr[i] * arr[i]
		}()
		tmp := <-ch
		res += tmp
	}
	fmt.Println(res)
}
