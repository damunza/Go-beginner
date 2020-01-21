package main

import "fmt"
func main(){
	fmt.Println("Enter temperature in degrees farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output:= ((input-32) *5/9) // remember that divisions and multiplications dont need brackets
	fmt.Println(output)

}