package main
import (
	"fmt"
	"strings"  // import  packages 
)

func  main() {

	fmt.Println("Hello string ")
	fmt.Println(strings.Compare("hii","hello"))  
	/* if a == b then output is 0
	   if  a  < b then -1
	   if  a  >  b then 1 
	*/
	fmt.Println(strings.Contains("hello world","llo"))
	fmt.Println(strings.Contains("helloworld","ok"))
	fmt.Println(strings.Contains("helloworld",""))
	fmt.Println(strings.Contains("",""))
// counting  strings
	fmt.Println(strings.Count("hello world","l"))

	// replace the strings 
	fmt.Println(strings.Replace("hello world he is no he","he","HE",2))
	fmt.Println(strings.Replace("okkk dkfjd kkk sffsk kk","k","W",-1)) // can use replaceALL
	// spliting string 
	fmt.Println(strings.Split("hello world this is me"," "))
	// 
	
}