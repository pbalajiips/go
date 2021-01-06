package main

import "fmt"

type FullNameType func(string, string) string

type Employee struct {
	FirstName, LastName string
	FullName            FullNameType
}

func main() {
	e := Employee{
		FirstName: "Ross",
		LastName:  "Geller",
		FullName: func(firstName string, lastName string) string {
			return firstName + " " + lastName
		},
	}

	fmt.Println(e.FullName(e.FirstName, e.LastName))
}
