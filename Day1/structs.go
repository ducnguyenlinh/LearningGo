package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	fmt.Println(person{"Linh", 23})

	fmt.Println(&person{"Duc Linh", 9})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
