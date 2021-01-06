package main

import (
	"fmt"
	"reflect"
)

type Sayer interface {
	Say() string
}

type Cat struct{}

func (c Cat) Say() string { return "meow" }

type Dog struct{}

func (d Dog) Say() string { return "woof" }

func main() {
	c := Cat{}
	fmt.Println("Cat says:", c.Say())
	d := Dog{}
	fmt.Println("Dog says:", d.Say())
	animals := []Sayer{c, d}
	fmt.Println(animals)
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}
	fmt.Println(animals)
}
