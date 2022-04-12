package main
import (
	"fmt"
	"time"
	"os/exec"
)

func main() {
	var x int 
	var y int 
	fmt.Println("enter first number : ")
	fmt.Scan(&x)
	time.Sleep(2*time.Second)

	fmt.Println("enter second  number : ")
	fmt.Scan(&y)
	// eq 
	result1 := x == y 
	fmt.Println("both are equal ",result1)
	// gt 
	result2 := x > y 
	out,err := exec.Command("date").Output()
	fmt.Println("value of operatos ",result2)
	time.Sleep(2*time.Second)
	fmt.Printf("current command output is %s \n",out)
	time.Sleep(1*time.Second)
	fmt.Println("error is ",err)
	
}