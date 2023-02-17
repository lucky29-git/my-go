package main

import (
	"bufio"
	"fmt"
	"os"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("Input from user")
 	reader :=	bufio.NewReader(os.Stdin)
	inp , _ := reader.ReadString('\n') 
    fmt.Println("is : " , inp)
}


func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}