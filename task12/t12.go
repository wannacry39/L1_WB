package main

import "fmt"

func main() {
	strs := [5]string{"cat", "cat", "dog", "cat", "tree"}
	set := []string{}
	hashmap := make(map[string]bool)

	for _, val := range strs {
		_, ok := hashmap[val]
		if !ok {
			set = append(set, val)
			hashmap[val] = true
		}
	}
	fmt.Print(set)
}
