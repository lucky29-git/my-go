package main

import "fmt"

func main() {

	

	days := []string{"Sunday" , "Tuesday" , "Wednesday" , "Friday" , "Saturday"}
	// fmt.Println(days)

	// for d:=0 ; d < len(days) ; d++{
	// 	fmt.Println(days[d])
	// }

	// if to run loop for full range can use this

	for i:= range days{
		fmt.Println(days[i])
	}

	//for-each loop

	for index , day := range days{                      //   or use err , synatx      for _ , day
		fmt.Printf("for index %v value is %v \n" , index , day)
	}

	// while loop
	
	// randomVal :=1
	// for randomVal < 10{
	// 	fmt.Println(randomVal)
	// 	randomVal++
	// }

	// while loop with break 

	randomVal :=1
	for randomVal < 10{
        if randomVal ==5{
			break
		} 
		fmt.Println(randomVal)
		randomVal++
	}

	// while loop with continue   and goto statements
	for randomVal < 10{
        if randomVal ==5{
			randomVal++ // so that next time when condition is checked for running loop it passes it and loop runs continue after 5
			continue
		} 

		if randomVal==7{
			goto hahaha                      //  or randomVal++
//                                                  goto hahaha                 so it will run loop after 7 also
		}
		fmt.Println(randomVal)
		randomVal++
	}

	hahaha:                                    // label 
		fmt.Println("HEHEHEHEHEHEHEHE")


}