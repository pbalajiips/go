package main

import (
	"encoding/json"
	"fmt"
)


// Hotel ...
type Hotel struct {
	Name   string `json: Name`
	Place  string  `json: Place`
	NoBeds int `json: NoBeds`

}

func main()  {
	var H Hotel // we must create the object 
	a := Hotel{Name: "Hilton", Place: "Chennai", NoBeds: 12}  
    b,err := json.Marshal(a)  

	err = json.Unmarshal(b, &H) // b has []bytes which will be the input for 
	if err != nil{
		fmt.Println("Unmarshal error")
	}

	fmt.Println(H)
}