package main
import (
	"fmt"
	"time"
	"context"
)

type SlowFunction func (string)(string,error)
type WithContext func(context.Context,string)(string,error)

func Timeout(f SlowFunction) WithContext {
	return func(ctx context.Context,arg string)(string,error){
		chres := make(chan string)
		cherr := make(chan error)

		go func(){
			res,err := f(arg)
			chres <- res
			cherr <- err
		}()

		select {
		case res := <- chres:
			return res,nil
		case <- ctx.Done():
			return "",ctx.Err()
		}
	}
}

func Slow(input string) (string, error) {
	time.Sleep(2 * time.Second) // Simulating a slow operation
	return "Slow operation completed", nil
}

func main(){
	ctx := context.Background()
	ctxt,cancel := context.WithTimeout(ctx,1*time.Second)
	defer cancel()
	timeout := Timeout(Slow)
	res,err := timeout(ctxt,"some context")
	fmt.Println(res,err)
}