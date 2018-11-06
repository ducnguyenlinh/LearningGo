package main

import "fmt"

func intSeq() func()int {
	i := 1
	return func()int {
		i += 2
		return i
	}
}

func main()  {
	nextInt := intSeq()
	fmt.Println(nextInt())

	fmt.Println(nextInt()) // retun int

	fmt.Println(intSeq()) // return function
}
