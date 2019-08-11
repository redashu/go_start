package main
import (
	"fmt"
	"time"
	"strings"
)

func  main() {

	fmt.Println(time.Now().UTC())
	fmt.Println(strings.Count("hello all","l"))
	// array
	var  a[3] int
	a[0] = 20
	a[1] =  30
	a[2] = 50
	fmt.Println(a[0])
	fmt.Println(a[0]+a[1])

	// an other way of maths
	p := [4]int{3,4,55,11}
	fmt.Println(p[1:4])
	fmt.Println(len(p))
	fmt.Printf("type of p is %T \n",a)
	p[0] = 900  // replacing  values
	fmt.Println(p)
	x :=  100
	y :=  "hello"
	aa :=  []int{4,77,22}    // slice literals 
	fmt.Println(x)
	fmt.Printf("%T \n",y)
	fmt.Printf("type of %T",aa)
	
}