package main

import (
	"fmt"
)

func main() {
	// there are 3 types how we declair variable
	// first method
	var i int
	i = 42
	i = 20
	fmt.Println(i)
	// second method
	var j int = 100
	fmt.Println(j)
	fmt.Printf("value of j is %v and type of data is %T\n", j, j)
	// third method and here we don't need to define datatype
	k := 200
	fmt.Println(k)

}

