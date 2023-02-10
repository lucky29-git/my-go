package main

import "fmt"

func main() {

	fmt.Println("")

	var fruits [4]string                      //  compulsory to provide the size of array

	fruits[0] = "apple"
	fruits[1] = "mango"
	fruits[2] = "banana"
	fruits[3] = "dragon fruit"

	fmt.Println("Fruit List" , fruits)
	fmt.Println("Fruit List length is " , len(fruits))

	// var vegetables [3]string{"potato " , " beas" , " mushroom "}
	// fmt.Println(vegetables)
}