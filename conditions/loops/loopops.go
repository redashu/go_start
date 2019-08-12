package main
import (
	"fmt"
)

func main()  {
    sum := 0
	for i := 0 ; i < 10 ; i++  {
		sum += i 
	}
	fmt.Println(sum)

	var a[5][3] int   // Two D  array 
	for j := 0 ; j < 5 ;j++ {
		for k := 0 ; k < 3 ; k++ {
			a[j][k] = j+k
		}
	}
	fmt.Println(a)
	for l := 0 ; l < 5 ; l++ {
		for m := 1 ; m < 6 ; m++ {
			fmt.Println("*")
		}
	}
}

