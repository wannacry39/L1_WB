package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(Reverse(s))
}

func Reverse(s string) string {
	sl := strings.Fields(s)
	j := len(sl) - 1
	for i := 0; i < j; i++ {
		sl[i], sl[j] = sl[j], sl[i]
		j--
	}
	return strings.Join(sl, " ")
}
