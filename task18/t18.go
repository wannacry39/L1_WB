package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	i int64
}

func newCounter() *Counter {

	return &Counter{0}
}

func (c *Counter) inc() {
	c.i++
}

func (c *Counter) get_val() int64 {
	return c.i
}

func main() {
	counter := newCounter()
	wg := sync.WaitGroup{}
	// m := sync.RWMutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter.i, 1)
			// m.Lock()
			// counter.inc()
			// m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter.get_val())
}
