package main

import (
	"fmt"
	"time"
	"log"
	"context"
)

func main(){
    start := time.Now()
	ctx := context.Background()
	userID := 10
	data,err := fetchUserData(ctx,userID)
	if err!=nil{
		log.Fatal(err)
	} else{
		fmt.Println("result: ",data)
		fmt.Println("took: ",time.Since(start))
	}
	
}

type Response struct{
	value int
	err error
}

func fetchUserData(ctx context.Context,userID int)(int,error){
	ctx,cancel := context.WithTimeout(ctx,time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	go func(){
		val,err := fetchThirdPartyStuff()
		respch <- Response{
			value: val,
			err: err,
		}
	}()

	for {
		select {
		case <- ctx.Done():
			return 0,fmt.Errorf("Timed out !!")
		case resp := <- respch:
			return resp.value,resp.err
		}
	}
}

func fetchThirdPartyStuff()(int,error){
	time.Sleep(time.Millisecond*300)
	return 666,nil
}