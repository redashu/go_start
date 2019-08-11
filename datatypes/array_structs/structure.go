package main
import  (

	"fmt"
)
// structure is something that we can use to group data together 

// defining a structure type 
type  human struct  {

	Firstname string
	lastname string
	age int
}


	// another example 
	type employ struct {

		name , job  , country string
		salary  int
	}

// executing  main function 
func main()  {

	var p human // here human is already define as struct so p is struct type
	fmt.Println(p) // with zero values 

	p1 := human{"ashutoshh","Singh",28}
	fmt.Println("Human1 : ",p1)

	p2 := human{Firstname: "Hello world"}
	fmt.Println(p2)

	e1 := employ {

		name: "mukesh",
		job: "CTO",
		country:  "India",
		salary:  50000,		
	}

	fmt.Println("name of employ : ",e1.name)

}



