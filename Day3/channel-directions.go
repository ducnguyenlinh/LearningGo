package main

import "fmt"

func pinger(pings chan <-string, msg string)  { // pings channel only write
	pings <- msg
}

func ponger(pings <-chan string, pongs chan<- string)  { // pings channel only write; pongs channel only read
	msg := <-pings
	pongs <- msg
}

func main()  {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	pinger(pings, "Message")
	ponger(pings, pongs)

	fmt.Println(<- pongs)
}
