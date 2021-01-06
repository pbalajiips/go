package main

import "fmt"

type Car struct {
	Brand    string
	Power    string
	Capacity int
}

// without pointer it will not modify the value  on the real object
// func (c Car) ChangePower(newpower string) {
// 	c.Power = newpower
// 	return
// }

func (c *Car) ChangePower(newpower string) {
	c.Power = newpower
	return
}

func main() {
	car := Car{
		Brand:    "Toyota",
		Power:    "2000CC",
		Capacity: 8,
	}
	fmt.Println("The Car's Capacity is ", car.Power)
	car.ChangePower("2500cc")
	//fmt.Println(car)
	fmt.Println("Now the car's Capacity is", car.Power)

}
