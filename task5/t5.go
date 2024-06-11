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
		if ctx.Err() == context.DeadlineExceeded {
			break
		}
		ch <- i
	}
}

func read(ctx context.Context, ch chan int) {
	for {
		if ctx.Err() == context.DeadlineExceeded {
			break
		}
		fmt.Println(<-ch)
	}
}
