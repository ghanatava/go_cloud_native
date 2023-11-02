package main

import "fmt"

func add(a int,b int) int{
	return a + b 
}

func factorial(n int) int{
	if n == 0 || n==1{
		return 1
	}
	return n*factorial(n-1)
}

func main(){
	var a,b int
	fmt.Println("Enter value of a")
	fmt.Scanf("%d",&a)
	fmt.Println("Enter value of b")
	fmt.Scanf("%d",&b)
	fmt.Println(add(a,b))

    fmt.Println("input the number for factorial")
	var n int
	_,err := fmt.Scanf("%d",&n)
	if err==nil{
		fact := factorial(n)
    	fmt.Println(fact)
	} else{
		fmt.Println(err)
	}
}