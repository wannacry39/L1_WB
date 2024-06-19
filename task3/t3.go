package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum int64
	arr := [5]int64{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < len(arr); i++ {
		go func(i int) {
			defer wg.Done()
			atomic.AddInt64(&sum, arr[i]*arr[i])
		}(i)
	}
	wg.Wait()
	fmt.Printf("%d\n", sum)

}
