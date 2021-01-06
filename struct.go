package main

import "fmt"

type Hotel struct {
	City           string
	Place          string
	Number_of_room int
	Open           bool
}

func main() {
	HolidayInn := Hotel{City: "scl", Place: "LuZ", Number_of_room: 10, Open: true}
	HolidayInn = Hotel{City: "scl", Place: "LuZ", Number_of_room: 10, Open: false}
	// Holiday.Name = "SCL"
	// Holiday.Place = "Luz"
	// Holiday.Number_of_room = 200
	fmt.Println(HolidayInn)

	//Anonymous struct can be used when we dont want to reuse the the struct.
	India := struct {
		Population  int
		No_of_state int
	}{
		Population:  100,
		No_of_state: 28,
	}
	fmt.Println(India)

	//Assign struct's Pointer using & operator and deferencing using * operator
	ITC := &Hotel{"Chennai", "Guindy", 2000, true}

	fmt.Println(ITC.City)

}
