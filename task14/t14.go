package main

import (
	"fmt"
)

func main() {
	var in interface{}
	in = 5
	typeof(in)
	in = "hello"
	typeof(in)
	in = true
	typeof(in)
	in = make(chan interface{})
	typeof(in)

}

func typeof(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("channel")
	}
}
