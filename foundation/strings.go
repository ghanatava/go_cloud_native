package main
import "fmt"

func main(){
    str := "foo"
	bytes := []byte(str)
	runes := []rune(str) 
	fmt.Printf("%7T %v\n",str,str)
	fmt.Printf("%7T %v\n",bytes,bytes)
	fmt.Printf("%7T %v\n",runes,runes)

}