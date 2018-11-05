package main

import "fmt"

func main()  {
	s := make([]string, 3)
	fmt.Println("Slice: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])

	fmt.Println("length: ", len(s))

	//append
	s = append(s, "d", "e")
	fmt.Println("append: ", s)

	//copy
	new := make([]string, 3)
	copy(new, s)               // copy from s to new
	fmt.Println(new)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:4]
	fmt.Println("sl2: ", l)

	l = s[1:]
	fmt.Println("sl3: ", l)

	twoD := make([][]int, 3)
	for i:=0; i<3 ; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j:=0; j<innerLen ; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
