package main

import "fmt"

func main() {

	//   their is No inheritence , concepts of super or parent classes in golang

	Lucky := User{"Lucky ", "lucky@dev.com", 20, true}
	fmt.Println(Lucky)
	fmt.Printf("Deatils of %v are : %+v \n" , Lucky.Name , Lucky)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
