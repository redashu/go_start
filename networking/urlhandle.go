package main

import (
	"fmt"
	"net/url"
)

func url1() {
	fmt.Println(myurl)
	// parsing 
	result , err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme) // getting protocol 
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

} 