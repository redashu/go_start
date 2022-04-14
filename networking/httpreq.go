package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func req1() {
	response,err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("type of response %T\n",response)
	defer response.Body.Close() // caller responsibity 
	// reading 
	databytes , err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)
}