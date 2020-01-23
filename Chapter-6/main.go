package main

// complex functions 
import "fmt"
func main(){
	basicNest()
// 	variadic()
// 	closure()
}

// an example of a Basic Nest function
func average(list [] float64) float64{
	/*
	a function that takes in an array does an action on it
	returns a float value
	*/
	total := 0.0
	for _, item := range list{
		total += item
	}
	return total/float64(len(list))
}

func basicNest(){
	/*
	this function holds the variables and the action done by 
	the average function
	*/
	x := [] float64{
		90,
		92,
		82,
		86,
		78,
	}
	fmt.Println(average(x))
}