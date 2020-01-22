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
	x["one"] = 1
	x["two"] = 2
	x["three"] = 3
	x["four"] = 4
	x["five"] = 5
	x["six"] = 6

/*
the following line prevents the calling of keys that do not
exist and receiving a 0 or ""
*/
	if name, ok := x["six"]; ok{
		fmt.Println( name, ok)
	}
}

