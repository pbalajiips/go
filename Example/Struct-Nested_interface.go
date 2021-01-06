package main

import "fmt"

type SpeedCal interface {
	Calculate() int
}

type Data struct {
	Network_Speed int
	HW_Speed      int
}

func (D Data) Calculate() int {
	return D.Network_Speed * D.HW_Speed
}

type App struct {
	Name  string
	Speed SpeedCal
}

func main() {

	Gmail := App{
		Name:  "Gmail",
		Speed: Data{10, 30},
	}
	fmt.Println(Gmail.Speed.Calculate())
}
