package main

import "fmt"

func zeroval(ival int)  {
	ival = 0
}

func zeroptr(iptr *int)  {
	*iptr = 0
}

func main()  {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)

	//-------------------------------->
	var x,z *int
	var y int
	y = 5
	x = &y
	z = &y

	fmt.Println("x = ", x)
	fmt.Println("*x = ", *x)

	*x = 1
	fmt.Println("new value: ", y)
	fmt.Println(*z)

	//---------------------------------->
	iptr := new(int) // create pointer by new()
	zeroptr(iptr)
	fmt.Println("new(): ", *iptr)
}