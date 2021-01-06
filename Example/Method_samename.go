package main

import "fmt"

type Walking struct {
	Speed    int
	Distance int
}

type Driving struct {
	Speed    int
	Distance int
}

func (w Walking) Calc() int {
	return w.Speed * w.Distance
}

func (d Driving) Calc() int {
	return d.Speed * d.Distance
}

func main() {
	man := Walking{10, 10}
	car := Driving{60, 20}

	fmt.Println("Man Walk in ", man.Calc())
	fmt.Println("Car max speed is ", car.Calc())
}
