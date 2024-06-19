package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Calling sleep func...")
	sleep(3 * time.Second)
	fmt.Println("End of main")
}

func sleep(n time.Duration) {
	timer := time.NewTimer(n)
	<-timer.C

}
