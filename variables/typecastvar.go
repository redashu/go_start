package main

import "fmt"

func main() {
	fmt.Println("Hello GO !")
	var x int = 200
	var y float32
	y = float32(x)
	fmt.Printf("hello conversion %v and type is %T \n", y, y)

}
