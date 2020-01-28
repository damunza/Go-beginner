package main

import (
	"fmt"
)

var x []float64

func main(){
	x :=[]float64{
		1.23,
		1.34,
		1.45, 
		1.67}
	y := x[1:3]
	fmt.Println(maxValue(y))
	fmt.Println(minValue(y))
}

// func createSlice(a, b int) []float64{
// 	/*
// 	a function to create slices out of any array provided starting from a and ending at b-1
// 	*/
// 	return x[a : b]
// }

func maxValue( y []float64) float64{
	/*
	function to return the maximum value in a slice
	*/
	var maxvalue float64
	// for x in the slice y 
	for _, x := range y{
		// starting from a maxvalue of 0  check if x is greater than maxvalue and make maxvalue x
		if maxvalue ==0 || x > maxvalue{
			// remember you are equating to x not assigning to x
			maxvalue = x
		}
		
	}
	return maxvalue
}

func minValue(y []float64) float64{
	/*
	function to return the minimum value of a slice
	*/
	var minvalue float64
	for _, x:= range y{
		if minvalue == 0 || x < minvalue{
			minvalue = x
		}
		
	}
	return minvalue
}


