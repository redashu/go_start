package main

import (
	"fmt"
	"time"
)

func main() {
	go task1("hello") // goroutine
	task1("world")
}
func task1(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(5*time.Millisecond)
		fmt.Println(s)

	}
}