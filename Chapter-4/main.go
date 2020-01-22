package main

import "fmt"
func main(){
	lessTen()
	multiplesThree()
	fizzBuzz()
}

/* unlike in python loops are bound by curly 
bracets not indenteds
*/

func lessTen(){
	for i:=1; i<=10; i++{
		fmt.Println(i)
	}
}

func multiplesThree(){
	for i:=1; i <=100; i++{
	if i % 3 == 0{
		fmt.Println(i)
	}
}
}

func fizzBuzz(){
	for i:=1; i <= 100; i++{
		if i % 3 == 0{
			fmt.Println("Fizz")
		}
		if i % 5 ==0{
			fmt.Println(" Buzz")
		}
		if i %3 != 0 && i%5 != 0{
			fmt.Println(i)	
		}
	}
}