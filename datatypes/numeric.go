package main
import  (

	"fmt"
"time"
//	"math"
)
func  main()  {
	var x int = 10  // basic type 
	j := x  //  example of interface type variable 
	var y float64 = 88.9
	fmt.Print(x,"\n")
	fmt.Println(j)
	fmt.Println(y)
	fmt.Println("Current time is ",time.Now().UTC())
	fmt.Println("please enter a number :- ")
	var a int 
	fmt.Scan(&a)
	fmt.Printf("data type of number is %T \n",a)  // here println will not work
}