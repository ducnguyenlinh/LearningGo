package main

import "fmt"

func first()  {
	fmt.Println("first")
}

func second()  {
	fmt.Println("second")
}

func division(a int, b int)int  {
	if b == 0 {
		defer func() {
			str := recover()
			fmt.Println(str)
		}()
		panic("Error because b = 0")
	}
	return a/b
}

func main()  {
	defer second()
	first()

	// first() --> second()
	fmt.Println(division(10,0))
}
