package main

import (
	"fmt"
	"net/url"
)

const myurl string ="https://lco.dev:3000/learn?coursename=reactjs&paymentid=arfaafaf" 
func main() {

	// We have url and extracting info from it 

	fmt.Println(myurl)
	
	// Parsing   --> means extracting some information from url

	result , _ :=  url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf(" \n The type of value of qparams %T \n" , qparams)

	fmt.Println(qparams["coursename"])

	// To print all key value pairs 
	for _ ,val := range qparams{
		fmt.Println(" The value is " , val)
	}

	// Now we have all the information and has to construct the url outofit 

	partsOfUrl := &url.URL{

		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawQuery: "user=lucky",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}