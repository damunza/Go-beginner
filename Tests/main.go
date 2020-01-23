package main

import "fmt"
func main(){
	halfType()
}

func halfType(){
	fmt.Println("Pick a number")
	var input int
	fmt.Scanf("%f", &input)
	output := input/2
	if input % 2 == 0{
		x := "True"
		fmt.Println(output, x)
	}else{
		x := "False"
		fmt.Println(output, x)
	}

}