package main

import "fmt"

func main() {

	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	res := make(map[int][]float32)

	for _, val := range arr {
		key := int(val/10) * 10
		_, ok := res[key]
		if !ok {
			res[key] = []float32{val}
		} else {
			res[key] = append(res[key], val)
		}

	}
	fmt.Println(res)

}
