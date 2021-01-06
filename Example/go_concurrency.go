package main

import (
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		// fmt.Println("in the loop")
		// amt := time.Duration(rand.Intn(20))
		// time.Sleep(time.Minute * amt)
		// fmt.Println("below  the loop")
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	fmt.Println("test")
	var input string
	fmt.Scanln(&input)
}
