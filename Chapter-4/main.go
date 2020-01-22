package main

import "fmt"
func main(){
	lessTen()
}

/* unlike in python loops are bound by curly 
bracets not indenteds
*/

func lessTen(){
	for i:=1; i<=10; i++{
		fmt.Println(i)
	}
}