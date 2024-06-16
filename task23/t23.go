package main

import "fmt"

func main() {
	var id int
	arr := []int{}
	res := []int{}
	for i := 0; i < 100; i++ {
		arr = append(arr, i)
	}

	fmt.Print("id of item: ")
	fmt.Scan(&id)
	res = arr[:id]
	arr = arr[id+1:]
	res = append(res, arr...)
	fmt.Println(res)
}
