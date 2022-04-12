package main

import (
	"fmt"
)

func main() {

	//fmt.Print("Hello world")
	// variable define 
	var Company_Name = "The xyz company"
	// constant for fixed value
	const exam_coupons = 50 
	var remaining_coupons = 50 
	//var  remaining_coupons = 50 
	fmt.Printf("company name is having %T, exam coupons is having %T ",Company_Name,exam_coupons)
	fmt.Println(Company_Name,"is having only",exam_coupons,"to offer")
	fmt.Printf("we have only %v Exam coupons and out of them only %v are left ", exam_coupons, remaining_coupons)	
	// better way to use variables
	fmt.Printf("welcomt to %v \n",Company_Name)
	// data types 
	var userName string
	var examBook int
	
	userName = "jack"
	examBook = 2 
	fmt.Printf("User %v purchased  %v exams \n",userName,examBook)

	
}
