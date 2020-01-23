package main

import ("fmt"; "math")
func main(){
	
}

type Cylinder struct{ 
	t float64 
	h float64 
	r float64
}

func surfaceArea(c *Cylinder) float64{
	rect := math.Pi * 2 * c.r * c.h
	circle :=math.Pi * c.r * c.r * c.t

	return rect + circle
}

func sample(){
	c := &Cylinder{1,14,7}
	fmt.Println(surfaceArea(c))
}
