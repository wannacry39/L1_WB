package main

import "fmt"

func main() {
	var str string
	fmt.Scanln(&str)
	rune_arr := []rune(str)
	j := 0
	for i := len(rune_arr) - 1; i > j; i-- {
		rune_arr[i], rune_arr[j] = rune_arr[j], rune_arr[i]
		j++
	}
	fmt.Println(string(rune_arr))

}
