package main
import  (
	"fmt"
)

func main()  {
	var x int
	fmt.Printf("Enter a number : ")
	fmt.Scan(&x)

	if x == 10 {
		fmt.Println("hello ")

	} else if  x < 10 {

		fmt.Println("world ")
	} else {
		fmt.Println("nothing is permanent ")
	}

}

