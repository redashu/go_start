package main

import (
	"fmt"
	"time"
)

var (
	x int    = 10
	y int    = 3
	z string = "hello"
)

func main() {
	fmt.Println("Starting with operators please wait..")
	time.Sleep(3 * time.Second) // holding for 3 seconds
	// calling function
	arithmetic()
	relational()
	logical()
}

func arithmetic() {
	var result1 int = x + y
	result2 := x - y
	fmt.Printf("addition is %v: \n", result1)
	fmt.Printf("substraction is %v \n", result2)
}

func relational() {
	eq := x == y
	fmt.Println("Result will be : ", eq)

	ne := x != y
	fmt.Println("Result will be : ", ne)

	lt := x < y
	fmt.Println("Result will be : ", lt)

	gt := x > y
	fmt.Println("Result will be : ", gt)

}

func logical() {
	if x != y && x <= y {
		fmt.Println("don't know what is happening !!")
	}
}

