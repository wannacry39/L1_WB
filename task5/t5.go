package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	var N int
	fmt.Scanln(&N)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(N))

	defer cancel()

	go read(ctx, ch)

	for i := 1; ; i++ {
		time.Sleep(300 * time.Millisecond)
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			ch <- i

		}

	}
}

func read(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			fmt.Println("time is over")
			return
		default:
			fmt.Println(<-ch)
		}
	}

}
