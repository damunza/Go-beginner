package main

import "fmt"
func main(){
	mean()
}

// if you dont include the float the valuess will be rounded down

func mean(){
	var x [5]float64
	//when declaring numerical values that need a math operation type is declared as a float
	x[0]=90
	x[1]=96
	x[2]=80
	x[3]=84
	x[4]=78
	total:=0.0 // add the .0 to create a float automatically
	for i:=0; i<5; i++{
		total += x[i]
	}
	fmt.Println(total/5)
}