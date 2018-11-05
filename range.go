package main

import "fmt"

func main()  {
	// ARRAY
	nums := []int{1, 3, 5, 7}
	sum := 0

	//--------------------------------->
	
	for i:= 0; i<len(nums) ; i++ {
		sum += nums[i]
	}

	fmt.Println("Example 1: ", sum)
	
	//////////////////////////////////

	sum = 0
	for _, num := range nums  {
		sum += num
	}

	fmt.Println("Example 2: ", sum)

	//---------------------------------->

	//MAP
	maps := map[string]int{"banana": 1, "apple": 3}
	for k, v := range maps {
		fmt.Printf("%s -> %d \n", k, v) // key and value
	}

	for k := range maps {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
