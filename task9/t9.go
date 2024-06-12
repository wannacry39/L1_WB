package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var N int
	fmt.Print("Введите длинну массива: ")
	fmt.Scanln(&N)
	arr := make([]int, N)
	for i := range arr {
		arr[i] = i
	}

	go func() {
		for _, val := range arr {
			ch1 <- val
		}
		close(ch1)
	}()

	go func() {
		for {
			val, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- val * val
		}
		close(ch2)
	}()

	for {
		val, ok := <-ch2
		if !ok {
			break
		}
		fmt.Print(" ", val)
	}
	fmt.Print("\n")
}
