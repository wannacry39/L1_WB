package main

import "fmt"

func main() {
	arr1 := []int{4, 5, 88, 13, 23, 99, 100}
	arr2 := []int{45, 5, 13, 77, 456, 134, 1}
	ch := make(chan int)
	exit := make(chan struct{}, 2)
	table := make(map[int]int)

	go func() {
		for i, val := range arr1 {
			ch <- val
			if i == len(arr1)-1 {
				exit <- struct{}{}
			}

		}
	}()

	go func() {
		for i, val := range arr2 {
			ch <- val
			if i == len(arr2)-1 {
				exit <- struct{}{}
			}
		}
	}()

	go hashmap(table, ch)
	<-exit
	<-exit
	for k, v := range table {
		if v != 1 {
			fmt.Println(k)
		}
	}
}

func hashmap(t map[int]int, ch <-chan int) {
	for {
		val := <-ch
		_, ok := t[val]
		if ok {
			t[val] += 1
		} else {
			t[val] = 1
		}
	}
}
