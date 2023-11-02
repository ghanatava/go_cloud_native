package main

import "fmt"

func product(factors ...int) int{
	p := 1
	for _,v := range factors{
		p *= v
	}
	return p
}

func add(nums ...int) int{
	s := 0
	for _,v :=range nums{
		s += v
	}
	return s
}

func main(){
	fmt.Println(product(1,2,3,5,6))
	fmt.Println(add(3,4,2,1,0))

	m := []int{1,3,4,5}
	fmt.Println(product(m ...))
}