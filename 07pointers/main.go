package main

import "fmt"

func main() {

	fmt.Println("Let's learn about packages")

	// var ptr *int                   //     * infront of int/string etc means it is a pointer 
	// fmt.Println("Default value of pointer " , ptr)


	myNum := 23
	
	var ptr = &myNum               // & is for refrencing to that number

	fmt.Println("It gives address of var " , ptr)
	fmt.Println("It gives value of var " , *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is " , myNum)



}