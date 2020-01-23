package main

// complex functions 
import "fmt"
func main(){
	basicNest()
	variadic()
	closure()
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

func add(x ...int) int{
	/*
	this function takes input of multiple integers 
	returns and integer value after addition
	*/
	start:= 0
	for _, item := range x{
		start += item
	}
	return start
}

func variadic(){
	/*
	an example of a variadic function that takes in a 
	collection of integers and adds them up
	*/
	fmt.Println(add(1,2,3,4))

}

func makeEvenGenerator() func() int{
	/*
	The return must be FUNC()
	The value is unsigned
	this is a generator that when run generates a new value
	every time its run
	*/
	start := 0
	return func() (ret int){
		/*
		this function needs a return with the word RET then 
		the type of data to run
		*/
		ret = start
		start +=2
		return
	}

}

func closure(){
	/*
	this is a function that executes the generator
	*/
	next := makeEvenGenerator()

	fmt.Println(next())
	fmt.Println(next())
}