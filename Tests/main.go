package main

import "fmt"
func main(){
	swap()
	largest()
}

func change(x, y * int){
	*x, *y = *y, *x
}

func swap(){
	a, b := 1, 2
	change(&a, &b)
	fmt.Println(a,b)
}

func big(x ...int)int{
	var current_no int
	for _,item := range x {
		// starting from my own value 0 
		// item equal to 0 check if variable larger than mine
		// if so change mine to the larger number
		if current_no == 0 || item > current_no{
			/*
			or is uded because it returns the correct statement
			of the two .
			assigns current number 0 then checks if the position
			iterated is larger than the current number returns 
			the true statement of the two
			*/
			current_no = item
		}
		
	}
	return current_no
}

func largest(){
	fmt.Println(big(1,5,3,4,9))
}