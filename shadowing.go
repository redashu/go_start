package main

import (
	"fmt"
)

var x int = 100

func main() {
	fmt.Println("value of x is %v", x)
	var x int = 200
	fmt.Printf("value of x i.e %v will used the most inner one \n", x)
}
