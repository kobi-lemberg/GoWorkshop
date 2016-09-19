package main

import (
	"math"
	"fmt"
)

type Circle struct {
	X,Y float64
	Radius float64
}

type Rectangle struct {
	X1,Y1, X2,Y2 float64
}

type Shape interface {
	Move(dx,dy float64)
	Area() float64
}

func (r *Rectangle) Area() float64  {
	dx := math.Abs(r.X1-r.X2)
	dy := math.Abs(r.Y1-r.Y2)
	return (dx * dy)
}

func (r *Rectangle) Move(dx,dy float64)  {
	r.X1 +=dx
	r.X2 +=dx
	r.Y1 +=dy
	r.Y2 +=dy
}


func (c *Circle) Area() float64  {
	return (math.Pi * c.Radius * c.Radius)
}

func (c *Circle) Move(dx ,dy float64)   {
	c.X +=dx
	c.Y +=dy
}

func main() {
	c:= &Circle{1,2,10}
	fmt.Printf("%#v\n",c)
	fmt.Printf("Area = %f\n",c.Area())

	c2:= Circle{1,2,100}
	fmt.Printf("%#v\n",c2)
	fmt.Printf("Area = %f\n",c2.Area())

	c.Move(1,-17)
	fmt.Printf("c moved:  %#v\n",c)

	r:= &Rectangle{1,1,4,4}
	fmt.Printf("%#v\n",r)
	fmt.Printf("R Area = %f\n",r.Area())

	shapes := []Shape{c,r,&c2}
	for _, s := range shapes{
		fmt.Println(s.Area())
	}
}
