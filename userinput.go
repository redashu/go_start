package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var x int
	var y int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&x)
	fmt.Println("you entered : ", x)
	// generating seed
	rand.Seed(time.Now().Unix())
	y=rand.Intn(20)  // it will generate random positive number less than 20
	println("Random number is : ",y)

}
