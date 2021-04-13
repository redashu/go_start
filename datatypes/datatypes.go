package main

import (
	"fmt"
	"time"
)

var (
	x int = 10
	y int = 3
)

func main() {
	var x bool = true
	fmt.Printf("Hello value of x is %v\n", x)
	fmt.Println("Please wait..!!")
	time.Sleep(4 * time.Second)
	// creating variables
	a := 1 == 1
	b := 2 == 1
	fmt.Printf("%v and %T \n", a, a)
	fmt.Printf("%v and %T \n", b, b)
	// calling functions
	mathops()
	operatos()
}

func mathops() {

	fmt.Printf("sum of numbers is %d \n", x+y)
	fmt.Printf("multiplication of numbers are %d \n", x*y)
	fmt.Println(x, x+5)
}

func operatos() {

	fmt.Println("x and y :- ", x&y)
	fmt.Println("x OR y :- ", x|y) // here binary operation are happening
}

