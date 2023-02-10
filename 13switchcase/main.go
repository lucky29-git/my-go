package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(" Let's play LUDO ")

	rand.Seed(time.Now().UnixNano())                //  rand is from pkg math  for genrating random no. in Unix nano second
	diceNumber := rand.Intn(6) +1                       // as 6 (range) is exclusive	

	fmt.Println("Value of dice is :" , diceNumber)
    
	switch diceNumber{
	case 1 :
		fmt.Println("Dice value is 1 and you can open")
	
	case 2 :
		fmt.Println("Move 2 spot")
	
	case 3 :
		fmt.Println("Move 3 spot")
		fallthrough                                    //    if we get 3 then next case statement will also run
	case 4 :
		fmt.Println("Move 4 spot")
	
	case 5 :
		fmt.Println("Move 5 spot")
	
	case 6 :
		fmt.Println("Move 6 spot and roll dice again")
	
	default :
		fmt.Println("Kya hua bhai")
	}

}