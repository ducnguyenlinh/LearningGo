package main

import "fmt"

func main()  {
	quere := make(chan string, 3)

	quere <- "one"
	quere <- "two"
	quere <- "three"

	close(quere)

	for element := range quere{
		fmt.Println(element)
	}
}
