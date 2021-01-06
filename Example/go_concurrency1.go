package main
import (
	"fmt"
	"time"
)


func Pong(c chan string){
  for j:=0;j<=5;j++{
	c <- "Pong"
	msi := <- c
	fmt.Println("test",msi)
  }
}

func Prepare(c <- chan string)  {
	for i := 0; i <=3  ;i++ {
		c <- "Ping"
	}
}

func Print(c chan string)  {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main(){
	var c chan string = make(chan string)

    go Pong(c)
	go Prepare(c)
	go Print(c)

	fmt.Println("Testing")
	var input string
	fmt.Scanln(&input)
}