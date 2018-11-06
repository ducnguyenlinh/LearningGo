package main

import (
	"fmt"
	"time"
)

func main()  {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "from 1"
		time.Sleep(time.Second * 2)
	}()

	go func() {
		c2 <- "from 2"
		time.Sleep(time.Second * 3)
	}()

	go func() {
		for i := 0; i< 10; i++{
			select {
			case msg1 := <-c1:
				fmt.Println("select channel 1: ", msg1)
			case msg2 := <-c2:
				fmt.Println("select channel 2: ", msg2)
			default:
				fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
