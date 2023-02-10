package main

import "fmt"

//this is allow  i.e.  variable outside method
// var newvar = 35           // it shows error only bcoz current variable is not used by any function

// First letter of variable is capital that means declaring it as public variable

const Lettercap string = " constants are also allowed xD"
 
func main() {

	var username string = "Lucky"
	fmt.Println(username)
	fmt.Printf("Type of variable is : %T \n ", username)

	var isBoolean bool = true
	fmt.Println(isBoolean)
	fmt.Printf("Type of variable is : %T \n ", isBoolean)

	var number int = 69
	fmt.Println(number)
	fmt.Printf("Type of variable is : %T \n ", number)

	var floatnum float64 = 0.99
	fmt.Println(floatnum)
	fmt.Printf("Type of variable is : %T \n ", floatnum)

	// implicit type

	var website = "www.hi.com"
	fmt.Println(website)


	// no var type
	numberOfUsers := 3524600
	fmt.Println(numberOfUsers)

}

// outside method not allowed
//  othervar := 356855


 

