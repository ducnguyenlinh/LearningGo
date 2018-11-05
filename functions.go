package main

import "fmt"

func plus(a int, b int)  int{
	return a + b
}

func main()  {
	res := plus(2,4)
	fmt.Println("2 + 4 =",res)
}
