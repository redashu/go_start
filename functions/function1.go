package main

import (
	"fmt"
	"os"
	"runtime"
)

/*  this is entry point for go compiler and doest not take any input
parameter also doest not return anything
*/
var x int
func init() {
	fmt.Println("i am init function , i will execute before main")
	x = 100 
}
func main() {
	fmt.Println("hello golang .!! ")
	fmt.Println(x)
	// calling area function 
	fmt.Printf("area of rectange is : %d \n",area(14,20))
	var p int = 10 
	var q int = 20 
	fmt.Printf(" %d %d \n",p,q)
	// 
	swaping(p, q)
	fmt.Printf("value of p is %d and q is %d ",p,q)
	
}

func init() {
	fmt.Println("i am init function , can be used to define global variables")
	// if init() function is not getting mac os then main function is not going to 
	// run 
	if runtime.GOOS == "darwin" {
		fmt.Println("execute this code")
	} else {
		fmt.Println("don't run this code")
		os.Exit(1)
	}
}