package main

import "fmt"

type Chef struct {
	Name       string
	Experience int
}

type ParkingCount struct {
	Number int
}

type Hotel struct {
	Location     string
	No_of_room   int
	chef         Chef //nested struct
	ParkingCount      //promated filed
}

func main() {
	Hyatt := Hotel{
		Location:     "Chennai",
		No_of_room:   100,
		chef:         Chef{Name: "Damu", Experience: 40},
		ParkingCount: ParkingCount{300},
	}

	fmt.Println(Hyatt)
}
