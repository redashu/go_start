package main
import  (
	"fmt"
	"time"
)

func  main() {
	x := 0 
	for x == 0 {   // there is for keyword for while loop 

		fmt.Println("Hello while")
		time.Sleep(2*time.Second)  // delay for 2  seconds
		fmt.Println("nice day")
		time.Sleep(2*time.Minute)   //  delay for 2 minutes
	}
}