package main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func main() {
	fmt.Println("Hii")

	content := " This is for revison purposes"

	file , err := os.Create("./filexyz.txt")
	checkNilError(err)
    length , err := io.WriteString(file , content)
	checkNilError(err)

	fmt.Println("print length" , length)

	defer file.Close()

	readfile("./filexyz.txt")

}
func readfile(filename string){
	data,err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println(" " , data)
	
}

func checkNilError( err error){
	if err!= nil {
		panic(err)
	}
}