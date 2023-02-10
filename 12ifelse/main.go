package main

import "fmt"

func main() {

	

	loginCount:= 10
	var result string 

	if loginCount<10{
		result = "Regular user"
	} else if loginCount >10 {
		result = " Chutiya User"
	} else {
		result = " Exactly 10 login"
	} 

	fmt.Println(result)

	// no need to create var in advance in go

	if 9%2==0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num:= 3 ; num<10{
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is less than 10")
	}

	// if err != nil{

	// }
}