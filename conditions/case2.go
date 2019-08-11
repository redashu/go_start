package  main
import  (
	"fmt"
	"time"
//	"strings"
)

func  main() {

	fmt.Println("Hello world !! lets start")
	fmt.Println("Today this time : ",time.Now().UTC())
	// var string 
	var x string = "Hello data"
	y :=  "Hello ashutoshh"
	fmt.Println(len(x))
	fmt.Println(y)
	if len(x) > len(y) {
		fmt.Println("as i KNow already..!")
	} else {

		fmt.Println("just got chance ")
	}
}