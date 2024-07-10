package main

import (
	"fmt"
	"math"
)


type Shape interface{
	area() float64
	
}

func main() {
	var shape Shape
	shape = circle{radius: 10,} 
	fmt.Println("Area of circle :",shape.area())

	shape = rectangle{height: 10 , width: 20,} 
	fmt.Println("Area of rectangle :",shape.area())
}

type circle struct{
	radius float64
}

type rectangle struct{
	height float64
	width float64
}

func (details circle) area() float64{
	return math.Pi*details.radius*details.radius
}

func (details rectangle) area() float64{
	return details.height*details.width
}

