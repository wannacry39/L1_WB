package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}

	var WorkerCount int
	var d int

	fmt.Print("Enter count of workers: ")
	fmt.Scanf("%d\n", &WorkerCount)
	fmt.Print("Enter time of work: ")
	fmt.Scanf("%d\n", &d)
	timer := time.NewTimer(time.Duration(d) * time.Second)
	go gen(ch, timer)
	wg.Add(WorkerCount)

	WorkerPool(WorkerCount, ch, &wg)

	wg.Wait()
}

func gen(ch chan int, t *time.Timer) {

	for i := 1; ; i++ {
		time.Sleep(300 * time.Millisecond)
		select {
		case <-t.C:
			close(ch)
			fmt.Println("time is over")
		default:
			ch <- i
		}
	}
}

func WorkerPool(count int, ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		go func(ID int) {
			for {
				val, ok := <-ch
				if !ok {
					defer wg.Done()
					break
				}
				fmt.Println("Worker:", ID, "Value:", val)
			}
		}(i)
	}

}
