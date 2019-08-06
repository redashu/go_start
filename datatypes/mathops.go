package main
import (

	"fmt"
	"time"
	"math"
)
func main()  {
	var x float64
	fmt.Println("current time is ",time.Now().UTC())
	fmt.Printf("Enter a number :- ")
	fmt.Scan(&x)
	fmt.Println("square root of ",x)
	fmt.Printf("you have %g ",math.Sqrt(x))
	// math.Sqrt is only accepting float64 numbers

// Now basic maths
	var a int = 10
	var b int = 20
	var c int = a+b
	fmt.Println(c)
	

}