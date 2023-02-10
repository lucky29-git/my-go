package main

import "fmt"

func main() {

	fmt.Println(" Lets's Learn Maps (also known as hash table { key value pair})")


	languages := make(map[string]string)                   //(map[int]string   key : int    value : string

	languages["JS"] = "JavaScript"
	languages["J"] = "Java"
	languages["PY"] = "Python"
	languages["ML"] = "Machine Learning"

	fmt.Println("List of all languages " , languages)
	fmt.Println(" JS stands for " , languages["JS"])

	delete(languages , "ML")
	fmt.Println(languages)
	
	// loops

	for key , value := range languages{

		fmt.Printf("for value %v , value is %v \n", key , value )            //   %v  is placeholder for value 
	}


}