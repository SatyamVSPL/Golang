package main

import "fmt"

type shape interface{
	getArea() float64
}

type triangle struct{
	base float64
	height float64
}

type square struct{
	sideLength float64
}



func (details triangle)getArea() float64{
	return 0.5*details.base*details.height
}

func (details square) getArea() float64{
	return details.sideLength*details.sideLength
}

func printArea(s shape){
	fmt.Println(s.getArea())
}

func main() {

	var Shape shape
	
	Shape= triangle{base:10 , height: 20}
	printArea(Shape)

	Shape= square{sideLength: 20}
	printArea(Shape)
}
