package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of golang")


	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 "))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))       // have to give same format (date time and day it's in documentation)

	// createdDate := time.Date()   have to provide all arguments in ()

	createdDate := time.Date(2023 , time.January , 27 , 02, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	// use go build 
	// to build an exe file (app) from it
	// GOOS="linux" go build     // to create app for linux system
	// GOOS="windows" go build     // to create app for windows system from linux

}