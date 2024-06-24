package main

import (
	"fmt"
	"reflect"
)

func main() {
	var vals []interface{}
	ch := make(chan int)
	vals = []interface{}{"word", 4, true, 4.56, []int{1, 2, 3}, ch}
	for _, val := range vals {
		fmt.Println(reflect.TypeOf(val))
	}

}
