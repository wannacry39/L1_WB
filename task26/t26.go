package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)

	fmt.Println(isUnique(str))
}

func isUnique(str string) bool {
	table := make(map[string]int)
	for _, val := range str {
		_, ok := table[strings.ToLower(string(val))]
		if ok {
			return false
		} else {
			table[strings.ToLower(string(val))] = 1
		}
	}
	return true
}
