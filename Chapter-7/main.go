package main

import ("fmt"; "math")
func main(){
	sample()
	sArea()
}

type Cylinder struct{ 
	/*
	a cylinder with the number of tops
	the height
	the radius
	*/
	t float64 
	h float64 
	r float64
}
type Rectangle struct{
	x1,x2,y1,y2 float64
}
// distance calculator
func distance(x,y float64) float64{
	a := x - y

	return a
}

// using pointers for the descriptive function and destination tags on the executed function

func surfaceArea(c *Cylinder) float64{
	rect := math.Pi * 2 * c.r * c.h
	circle :=math.Pi * c.r * c.r * c.t

	return rect + circle
}

func sample(){
	c := &Cylinder{1,14,7}
	fmt.Println(surfaceArea(c))
}

// using pointers in the executed function and a calling the descriptive function like a method

func (r *Rectangle) area() float64{
	/*
	this function has a receiver called area
	methods do not have names just receivers
	*/

	w := distance(r.y2,r.x1)
	l := distance(r.x2,r.x1)

	return l * w
}

func sArea(){
	/*
	when declaring variables using a type use curly brackets
	*/
	a := Rectangle{0,4,0,3}
	fmt.Println(a.area())
}
/**/