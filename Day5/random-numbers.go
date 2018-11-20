package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(20))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Println((rand.Float64()*5) + 10)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100))

	s2 := rand.NewSource(90)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100))
}
