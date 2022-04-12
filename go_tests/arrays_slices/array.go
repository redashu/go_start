package main

import (
	"fmt"
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
	
	//var exam_users = [10] string {"ashu","mayank","ujwal"}
	var exam_users [10] string
	exam_users[0] = "ashu"
	exam_users[1] = "mayank"
	exam_users[2] = "ujwal"

	// variables 
	var userName string
	var lastName string
	var array_use [5] string

	// userinput 
	fmt.Println("Enter First name : ")
	fmt.Scan(&userName)

	fmt.Println("Enter last name : ")
	fmt.Scan(&lastName)

	array_use[0] = userName + " " + lastName

	fmt.Printf("whole array %v and %v \n",exam_users,array_use)
	fmt.Printf("first value of array %v \n",exam_users[0])
	fmt.Printf("type of array : %T \n",exam_users)
	fmt.Printf("length of array %v \n",len(exam_users))
	
}
