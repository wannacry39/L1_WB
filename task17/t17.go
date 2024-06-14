package main

import "fmt"

func main() {
	arr := []int{1, 4, 8, 9, 11, 65, 103, 677, 900, 1024}

	res := BinarySearch(arr, 900)

	fmt.Println(res)
}

func BinarySearch(arr []int, n int) interface{} {
	left := 0
	right := len(arr) - 1
	for right >= left {
		mid := (right + left) / 2
		if arr[mid] == n {
			return mid
		}
		if arr[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return "No such value"
}
