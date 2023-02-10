package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Let's Learn slices")

	var fruitList = []string{"Apple","Mango"}

	fmt.Printf("Data type of fruitList : %T \n  i.e. slices \n" , fruitList )

	// To add more fruits in list
	fruitList = append(fruitList, "Banana" , "cherry" , "pineapple")
	fmt.Println(fruitList)

    fruitList = append(fruitList[1:3], )        // It will modify list such that it contains elements from index 1 to 3 (3 is exclusive)
	fmt.Println(fruitList)
	
	

	//  fruitList = append(fruitList[2:], )
	//  fmt.Println(fruitList)

	// fruitList = append(fruitList[:3], )
	// fmt.Println(fruitList)

	highscores := make([]int , 4)            // right now it is array as we are providing size

	highscores[0] = 101
	highscores[1] = 111
	highscores[2] = 222
	highscores[3] = 333
	// highscores[4] = 444

	fmt.Println(highscores)
//                                                slices is nothing but an array with more funtionability
	highscores = append(highscores, 777,888,55,66,4,2,45,56 )    // append wll reallocate the memory (here it is slices)
	fmt.Println(highscores)

	sort.Ints(highscores)           // sorts pkg sotrs the slice , Ints means in increasing order  (type sort.  to see more funs)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))        // is slice sorted ? 


	// Removing value from slices based on index 

	var courses = []string{"java" , "docker" , "golang" , "python" , "machine learning"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index] , courses[index+1 : ]... )       // first argument will take list from 0 to index(exclude) and then @nd argument take list from index+1 till end
    fmt.Println(courses)                 // goland gets deleted
}