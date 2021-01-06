package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	height float64
	weight float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.height * r.weight
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.weight + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

func main() {
	var s Shape
	s = Rect{3.0, 5.0}
	r := Rect{3.0, 5.0}

	fmt.Printf("the type of in s is %T\n", s)
	fmt.Printf("the type of in r is %T\n", r)
	fmt.Println("the value calculated from Interface", s.Area())
	fmt.Println("the value calculated from Struct", r.Area())
	fmt.Println("s == r is", s == r)
	//fmt.Printf("%T", s)
	s = Circle{3.0}
	//r = Circle{3.0} // we cannot use Circle literal (type Circle) as type Rect in assignment. we cannot use struct object to create another object but with ibterface we can do until methods
	fmt.Printf("the type of in s is %T\n", s)
	fmt.Printf("the type of in r is %T\n", r)
	fmt.Println("the value calculated from Interface", s.Area())
	fmt.Println("the value calculated from Struct", r.Area())
	fmt.Println("s == r is", s == r)

}
