package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x int = 100
	var y string
	y = strconv.Itoa(x)
	fmt.Printf("value is %v and type is %T \n", y, y)
}

