package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal("invalid input")
	}
	arr := strings.Fields(str)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	str = strings.Join(arr, " ")
	fmt.Println(str)
}
