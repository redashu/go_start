package main

import (
	"fmt"
)

var x int = 100 // package level variable
var Y int = 500 // if uppercase var name then compiler expose it
// to outside world
func main() {
	fmt.Println("value of x is %v", x)
	var x int = 200 // function level variable
	var a, b int = 59, 68
	fmt.Printf("value of x i.e %v will used the most inner one \n", x)
	fmt.Println("variable values are %v and %v ", a, b)
	// calling check function
	check()
}

func check() {
	fmt.Printf("here is the thing %d \n", x)
}
