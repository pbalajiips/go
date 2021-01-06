package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["Name"] = "Balaji"
	m["Nation"] = "India"
	fmt.Println("My name is ", m["Name"])
}
