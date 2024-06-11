package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("End of main")
	runtime_goexit()
	ctx_ending()
	chan_stop()
}

// Goroutine shutdown by using runtime.goexit

func runtime_goexit() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("runtime_goexit() started")
		defer fmt.Println("runtime.Goexit()")
		defer wg.Done()
		for i := 0; ; i++ {
			time.Sleep(300 * time.Millisecond)
			fmt.Println(">>>")
			if i == 5 {
				runtime.Goexit()
			}
		}
	}()
	wg.Wait()
}

// Goroutine shutdown with context

func ctx_ending() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context, ch chan string) {
		defer wg.Done()
		defer fmt.Println("grtn dying...")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx canceled")
				return
			case val := <-ch:
				time.Sleep(1 * time.Second)
				fmt.Println(val)
				cancel()
			}
		}

	}(ctx, ch)
	ch <- "ctx_ending is working"
	wg.Wait()
}

// Goroutine shutdown by sending signal in the channel
func chan_stop() {
	wg := sync.WaitGroup{}
	ch := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("chan_stop dying...")
		for {
			select {
			case <-ch:
				fmt.Println("Stop signal")
				return
			default:
				time.Sleep(300 * time.Millisecond)
				fmt.Println("chan_stop working...")

			}
		}
	}()
	time.Sleep(5 * time.Second)
	ch <- true
	wg.Wait()
}
