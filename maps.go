package main

import "fmt"

func main()  {
	m := make(map[string]int)

	m["k1"] = 5
	m["k2"] = 7

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("value 1: ", v1)

	delete(m, "k2")
	fmt.Println("delete: ", m)

	value_null, prs := m["k3"]
	fmt.Println(value_null, prs)

	_, prsss := m["k4"]
	fmt.Println("value null: ", prsss)

	n := map[string]int{"foo1": 2, "foo2": 5 }
	fmt.Println(n)
}
