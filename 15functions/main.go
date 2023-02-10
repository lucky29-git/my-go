package main

import "fmt"

func main() {
	fmt.Println(" Welcome to Main Function")

	greetings()

    
	// function inside function is not allowed

	// func greetingsTwo(){
	// 	fmt.Println("Good night")
	// }
	// greetingsTwo()

	result := adder(3,5)
	fmt.Println(result)

	resultX := proAdder(3,3,4,56,7)
	fmt.Println(resultX)

	proresult , myMessage := proAdder2(6,8,9,5,21,5) 
	fmt.Println(proresult)
	fmt.Println(myMessage)
}

func adder(valOne int, valtwo int) int {
	return valOne + valtwo
}

func proAdder (values ... int) int{
	total := 0

	for _ , val := range values{
		total += val ;
	}
	return total 
}

func proAdder2 (values ... int) (int,string){
	total := 0

	for _ , val := range values{
		total += val ;
	}
	return total , " the result is infront of you"
}

func greetings(){
	fmt.Println("Good morning")
}











