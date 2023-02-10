package main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func main() {

	fmt.Println("")

	content:= " This is going to be in a file"

	file , err :=os.Create("./fileabc.txt")
	if err != nil {
		panic(err)
	}

	length , err := io.WriteString(file , content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length of file is " , length)
	defer file.Close()
    readfile("./fileabc.txt")
}

func readfile( filename string)  {	
 	data , err := ioutil.ReadFile(filename)
 	checkNilError(err)
 	// fmt.Println("When Data is readed in file it reads in byte format not in string format \n" , data)
  fmt.Println("But convert it in string format so that we can read \n" , string(data))
}

func checkNilError( err error)  {
	if err != nil {
		panic(err)
	}
}