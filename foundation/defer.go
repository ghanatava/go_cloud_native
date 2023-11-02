package main

import (
	"fmt"
	"os"
)

func main(){
	file,err := os.Create("/tmp/foo.txt")
	defer closeFile(file)
	if err!=nil{
		return
	}

	_,err = fmt.Fprintln(file,"your mother was a hamster")
	if err != nil{
		return
	}

	fmt.Println("File written successfully")

}

func closeFile(f *os.File){
	if err := f.Close(); err != nil{
		fmt.Println("error closing file:",err.Error())
	} else{
		fmt.Println("file closed successfully")
	}
}