package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}
	go gen(ch)
	var WorkerCount int
	fmt.Scanf("%d\n", &WorkerCount)

	wg.Add(WorkerCount)

	WorkerPool(WorkerCount, ch, &wg)

	wg.Wait()
}

func gen(ch chan int) {
	for i := 1; ; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}
}

func WorkerPool(count int, ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		go func(ID int) {
			for {
				val, ok := <-ch
				if !ok {
					wg.Done()
					break
				}
				fmt.Println("Worker:", ID, "Value:", val)
			}
		}(i)
	}

}
