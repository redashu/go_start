package main
import (
	"fmt"
	"time"
)

func main()  {

	fmt.Println("Hello world guys ..!")
	fmt.Println(time.Now().UTC())
	var x string = "ok"
	var  y  string 
	fmt.Printf("please enter a string :- ")
	fmt.Scan(&y)
	fmt.Println("i am saying ",x)
	fmt.Println("you have entered ",y)
	fmt.Println(x+" "+y)  // adding string
	fmt.Println(x," ",y)  // adding string 
	fmt.Println("length of string ",y," is ",len(y))  // ;ength of the string
	fmt.Println(string(y[1])) // string  indexing will be done
	fmt.Println(y[1]) //  this will print ascii code 
	fmt.Println("string range is ",string(y[0:4]))  // range in string 
	fmt.Println("")

	

}