package main

import "fmt"
func main(){
	swap()
}

func change(x, y * int){
	*x, *y = *y, *x
}

func swap(){
	a, b := 1, 2
	change(&a, &b)
	fmt.Println(a,b)
}