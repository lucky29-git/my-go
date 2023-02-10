package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const url ="https://lco.dev"
func main() {

	fmt.Println(" Web request")
	response , err :=http.Get(url)
	checkNilError(err)
	fmt.Printf("Type of response is %T \n " , response)

	
	data , err :=ioutil.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println(" data " , string(data))

	defer response.Body.Close()      // Imp

}

func checkNilError( err error)  {
	if err != nil {
		panic(err)
	}
}