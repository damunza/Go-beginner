package main

import "fmt"
func main(){
	/*
	this function is used to call 
	everything that should run when the program is launched 
	*/
	farenCel()
	meterFeet()

}

func farenCel(){
	fmt.Println("Enter temperature in degrees farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output:= ((input-32) *5/9) // remember that divisions and multiplications dont need brackets
	fmt.Println(output)

}

func meterFeet(){
	fmt.Println("Feet to convert to meters:")
	var input float64
	fmt.Scanf("%f", &input)

	output:= (input/0.3048)
	fmt.Println(output)
}