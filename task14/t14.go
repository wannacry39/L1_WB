package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val interface{}

	val = []string{"adas", "ghh"}

	fmt.Println(reflect.TypeOf(val))

}
