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

}