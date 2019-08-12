package  main
import  (
	"fmt"
	"time"

)

func  sayhello() {
	fmt.Println("Hello function !!")
	time.Sleep(3*time.Second)
}

func  add(x int , y int) int {  // here function name , argument and return type

	return  x + y 
}

func  sum(x,y int) int {
	return x*y
}

//  multiple  results return 
func  swap(x , y string) (string , string) {
	return y  , x

}

func  main()  {

	sayhello()
	fmt.Println(add(5,8))
	fmt.Println(sum(55,77))
	fmt.Println(swap("Hello","world"))
}