package main

import "fmt"

func main(){
    ch := make(chan string,2)
	ch <- "foo"
	ch <-"bar"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(ch)

}