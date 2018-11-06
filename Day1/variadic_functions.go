package main

import "fmt"

func total_func(nums ...int)  {
	fmt.Print(nums, " ")

	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("total: ", total)
	fmt.Println("num[2]: ", nums[2])
}

func main()  {
	total_func(1, 2, 3)

	nums := []int{1, 3, 5, 7, 9}
	total_func(nums...)
}
