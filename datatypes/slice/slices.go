// Golang program to illustrate how
// to create a slice using a slice
// literal
package main

import "fmt"

func main() {

	// Creating a slice
	// using the var keyword
	var my_slice_1 = []string{"hello", "for", "slice"}

	fmt.Println("My Slice 1:", my_slice_1)

	// Creating a slice
	//using shorthand declaration
	my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice 2:", my_slice_2)

	// using array 
	arr := [7]string{"This", "is", "the", "tutorial",
                         "of", "Go", "language"}
 
    // Display array
    fmt.Println("Array:", arr)
 
    // Creating a slice
    myslice := arr[1:6]
 
    // Display slice
    fmt.Println("Slice:", myslice)
 
    // Display length of the slice
    fmt.Printf("Length of the slice: %d", len(myslice))
 
    // Display the capacity of the slice
    fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}
