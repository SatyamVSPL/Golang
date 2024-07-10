package main

import (
	"fmt"
	"math"
)

type Shape interface{
	area() float64
}

type circle struct{
	radius float64
}

type rectangle struct{
	width float64
	height float64
} 

func main() {

	c := circle{radius:10,}
	r := rectangle{width: 10, height: 20,}
	
	// fmt.Println(c.area(),": This is area of circle\n",r.area(),": This is area of rectangle")
	fmt.Println(calculateShape(c))
	fmt.Println(calculateShape(r))

}

func (c circle) area()float64 {
	return math.Pi *c.radius *c.radius		
}

func (r rectangle) area()float64{
	return (r.width)*(r.height)
}

func calculateShape(details Shape) float64 {
	return details.area()
}