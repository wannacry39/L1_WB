package main

import (
	"fmt"
)

func main() {

	keys := []string{"first", "second", "third", "fourth", "fifth"}
	vals := []int{1, 2, 3, 4, 5}

	digits := make(map[string]int)
	key_ch := make(chan string)
	val_ch := make(chan int)

	go func() {
		for i := range keys {
			key_ch <- keys[i]
		}
		close(key_ch)
	}()

	go func() {
		for i := range vals {
			val_ch <- vals[i]
		}
		close(val_ch)
	}()

	for {
		key, ok := <-key_ch
		if !ok {
			break
		}
		val := <-val_ch

		digits[key] = val

	}

	fmt.Println(digits)
}
