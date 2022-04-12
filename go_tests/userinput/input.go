package main

import (
	"fmt"
	"time"
)

func main() {

	//fmt.Print("Hello world")
	// variable define 
	var Company_Name = "The xyz company"
	// constant for fixed value
	const exam_coupons int  = 50 
	var remaining_coupons int = 50 
	//var  remaining_coupons = 50 
	fmt.Printf("company name is having %T, exam coupons is having %T ",Company_Name,exam_coupons)
	fmt.Println(Company_Name,"is having only",exam_coupons,"to offer")
	// better way to use variables
	fmt.Printf("welcomt to %v we have %v \n",Company_Name,remaining_coupons)
	// data types 
	var userName string
	var lastName string
	var email string
	var examBook int
	// taking user input as username  
	fmt.Println("Enter your user name : ")
	fmt.Scan(&userName)
	// last name of user 
	fmt.Println("Enter last name : ")
	fmt.Scan(&lastName)
	// user email id 
	fmt.Println("Enter email id : ")
	fmt.Scan(&email)
	// here & means pointer 
	fmt.Println("Enter number of exam coupons : ")
	fmt.Scan(&examBook)
	fmt.Println(&examBook) // printing memory location of that variable 
	fmt.Printf("User %v %v purchased  %v exams \n",userName,lastName,examBook)
	time.Sleep(5*time.Second)
	fmt.Printf("Email will be sent at %v \n",email)

	// Remaining tickets 
	
	remaining_coupons = remaining_coupons - examBook
	fmt.Println("Remaining tickets are ",remaining_coupons)

	
}
