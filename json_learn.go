package main

import (
	"encoding/json"
	"fmt"
)

type Food struct {
	Style string
}

// Hotel ...
type Hotel struct {
	Name   string
	Place  string
	NoBeds int
	Food
}

func main() {

	H := Hotel{Name: "Hilton", Place: "Chennai", NoBeds: 12}
	J := Hotel{Name: "Jasmine", Food: Food{Style: "Chinese"}}
	data, err := json.Marshal(H)

	if err != nil {
		fmt.Println("test")
	}
	jdata, err := json.Marshal(J)

	fmt.Println(string(data))
	fmt.Println(string(jdata))

}
