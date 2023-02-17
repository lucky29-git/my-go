package main

import (
	"fmt"
)

func main() {

 fmt.Println("Hii my name is lucky ")

}


func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}