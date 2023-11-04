package main

import (
    "fmt"
	"math"
	"bufio"
)
type Shape interface{
	Area() float64
}

type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return math.Pi*c.Radius*c.Radius
}

type Rectangle struct{
	Width,Height float64	
}

func (r Rectangle) Area() float64{
	return r.Width*r.Height
}

func printArea(s Shape){
	fmt.Printf("%T's area is %0.2f\n",s,s.Area())

}


func main(){
    r := Rectangle{Width:5,Height:4}
	printArea(r)

	c := Circle{Radius:3.14}
	printArea(c)
    
	var s Shape
	s = Circle{}
	c1 := s.(Circle)
	fmt.Printf("%T\n",c1)

}