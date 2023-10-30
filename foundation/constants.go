package main

import "fmt"

const language="Go"
var favourite bool=true


func main(){
	const text = "Does %s rule ? %t!"
	output := fmt.Sprintf(text,language,favourite)
    fmt.Println(output)	

}