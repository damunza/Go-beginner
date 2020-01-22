package main

import "fmt"
func main(){
	// mean()
	dictionary()
}

// if you dont include the float the valuess will be rounded down

// func mean(){
// 	x := [5]float64{
// 		90,
// 		96,
// 		80,
// 		84,
// 		78,
// 	}
// 	/*when declaring numerical values that need a 
// 	math operation type is declared as a float
// 	*/
// 	total:=0.0 // add the .0 to create a float automatically
// 	for _, value := range x{
// 		total += value
// 	}
// 	fmt.Println(total/float64(len(x)))
// }

/*
- dealing with map< dictionaries> data types 
- make key word is used to create maps
*/
func dictionary(){
	x := make(map[string]int) // declaring key data type and value data type
	x["key"] = 10
	fmt.Println(x)
}