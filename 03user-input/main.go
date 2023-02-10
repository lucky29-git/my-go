package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome in user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating :")

	// comma ok syntax  or  comma error syntax

	input , _ := reader.ReadString('\n')                 //  or input , err := reader.ReadString('\n')
	//                                                       i.e.  if everything is correct while taking input then it will store in input var
	//                                                      and if anything goes wrong (suppose an error occur) then it save error in err var
	fmt.Println("Thanks for rating " , input)

	fmt.Printf("Type of rating is %T" , input)

}