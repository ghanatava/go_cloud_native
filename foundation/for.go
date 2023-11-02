package main

import "fmt"

func main(){
	s := []int {1,3,5,9,23}
	for i,v := range s { 
		fmt.Printf("%v -> %v\n",i,v)
	}
	
	for _,v := range s{
		fmt.Println(v)
	}

	//looping over maps

	freezing := make(map[string]float32)
    freezing["celsius"] = 0.0
	freezing["farhenhite"] = 32.0
	freezing["kelvin"] = 273.2

	for key,val := range freezing{
		fmt.Println(key,"->",val)
	}
}

