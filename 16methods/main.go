package main

import "fmt"

func main() {

	Lucky := User{"Lucky" , "afhnf" , 20 , true}

	Lucky.GetStatus()

}

func (u User) GetStatus() {
	fmt.Println(" Is user active " , u.Status)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}