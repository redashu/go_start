package main

import (
	"fmt"
	"os"
)
func main() {
	fmt.Println("getting started !!..")
	cmdin := os.Args[1:]
	for i,item := range cmdin {
		fmt.Println(i)
		fmt.Println(item)
	}
	fmt.Println(cmdin[0]) 
	if cmdin[0] == "hello" {
		fmt.Println("Hello world this is mac")
	} else {
		fmt.Println("Hello data : ")
	}
}
