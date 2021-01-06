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
	Name string
	SpeedCal
}

func main() {

	Gmail := App{
		Name:     "Gmail",
		SpeedCal: Data{10, 30},
	}
	fmt.Println(Gmail.Calculate())
	//only methods are promoted wtih anonymous function not the fileds
	//fmt.Println(Gmail.Network_Speed)

}
