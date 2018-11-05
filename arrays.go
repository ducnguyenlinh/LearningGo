package main

import "fmt"

func main()  {
	var a [5]int
	fmt.Println("array: ", a)

	a[3] =  100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[3])
	fmt.Println("length: ", len(a))

	b := [6]float64{11, 22, 33, 44, 55, 66}
	fmt.Println(b)
}
