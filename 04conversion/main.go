package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our Pizza shop")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')
	
	fmt.Println("Thanks for rating " , input)

	// Converting string to integer

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if (err != nil){
			fmt.Println(err)
	} else {
		fmt.Println("Added one in rating" , numRating + 1)
	}



}