package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{4, 1, 8, 6, 10, 1, 14}
	fmt.Println(quickSortStart(arr))

}
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func partition(arr []int, low, high int) ([]int, int, int) {
	pivot := arr[rand.Intn(high-low)+low]
	for low <= high {
		for ; arr[low] < pivot; low++ {
		}
		for ; arr[high] > pivot; high-- {
		}
		if low <= high {
			arr[low], arr[high] = arr[high], arr[low]
			low++
			high--
		}
	}

	return arr, low, high
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		arr, left, right := partition(arr, low, high)
		arr = quickSort(arr, left, high)
		arr = quickSort(arr, low, right)
	}
	return arr
}
