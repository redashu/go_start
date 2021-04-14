package main

import (
	"fmt"
	"time"
)

var (
	x string = "Hello World"
	y string = "Golang is the fastest"
)

func main() {
	fmt.Println("value of x is ", x)
	time.Sleep(2 * time.Second)
	fmt.Println("y is having \n value : ", y)
	fmt.Println(`Hello \n world`)
	fmt.Println("length of string ", len(x))
	fmt.Println("indexing of string ", y[0])
	fmt.Println("character pring ", string(y[0]))
	fmt.Println(string(y[0:4]))
	// Note :- y="GO"---> y[0] -- will print ascii value
	//  y="Go" -- > string(y[0]) -- will print first character
}

