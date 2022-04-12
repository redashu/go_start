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
	fmt.Println("Hello world with new line ")
	fmt.Println("welcome to ",Company_Name)
	fmt.Println(Company_Name," is offering you exam portal ")
	fmt.Println(Company_Name,"is having only",exam_coupons,"to offer")
	fmt.Printf("we have only %v Exam coupons and out of them only %v are left ", exam_coupons, remaining_coupons)	
	// better way to use variables
	fmt.Printf("welcomt to %v \n",Company_Name)
	// data types 
	
}
