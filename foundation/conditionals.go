package main

import (
	"fmt"
	"os"
)

func main(){
	if _,err := os.Open("foo.txt"); err != nil{
		fmt.Println("err!")
	} else{
		fmt.Println("all works well!")
	}

	//switch case 
        i := 0
        switch i%3{
            case 0:
                fmt.Println("Zero")
            case 1:
                fmt.Println("one")
            default:
                fmt.Println("Huh?")
        }	
}
