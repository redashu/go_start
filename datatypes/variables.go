package main

import  (
	"fmt"
	"time"
)
func main()  {
	fmt.Println("Hello world")
	time.Sleep(4*time.Second)

	var  x  int //  simple variable initialize
	var  y  float64 = 5.8  // float 

	fmt.Println(x,y)
/*
	var  (         
		b  string
		c bool
	)
*/
	var  v , e , z  =  3 ,5.5 ,"hello"
	fmt.Println(v,e,z) 

	v = 88 // replacing 
	fmt.Println(v)
}