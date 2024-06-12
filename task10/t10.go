package main

import "fmt"

func main() {

	arr := []float32{-40.4, 21.6, 5.5, -17.0, -3.9, 0}

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
