package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	createFile()
}
func createFile() {
	fmt.Println("writing file in golang ")
	file, err := os.Create("test.txt") // everytime create new file 
	if err != nil {
		log.Fatalf("failed creating file: %s",err)
	}
	defer file.Close() //closing file after file has been created 
	len , err := file.WriteString("welcome to golang."+
					"are adding string data")
	if err != nil {
		log.Fatalf("failed in writing file : %s \n",err)
	}	
	fmt.Printf("\n FIle Name: %s",file.Name())
	fmt.Printf("\n length in bytes %d ",len)

}