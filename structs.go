package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	fmt.Println(person{"Linh", 23})

	fmt.Println(&person{"Duc Linh", 9})
}

