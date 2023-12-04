package main

import (
	"fmt"
	"context"
	"time"
    "sync"
	"errors"
)

type Circuit func(ctx context.Context)(string,error)

func Breaker(circuit Circuit,failureThreshold uint)Circuit{
	var consecutiveFailure int
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context)(string,error){
		m.RLock()
		d := consecutiveFailure - int(failureThreshold)
		if d >=0 {
			shouldRetryAt := lastAttempt.Add(time.Second*2<<d)
			if !time.Now().After(shouldRetryAt){
				m.RUnlock()
				return "",errors.New("Service Unreachable")
			}
		}
		m.RUnlock()
		response , err := circuit(ctx)
		m.Lock()
		defer m.Unlock()

		lastAttempt=time.Now()

		if err!=nil{
			consecutiveFailure++
			return response,err
		}
		consecutiveFailure = 0
		return response,nil
	}
}

func SimulatedCircuit(ctx context.Context)(string,error){
	attempts := 3
	if attempts > 0{
		attempts--
		return "",errors.New("Simulated Response")
	}
	return "Success",nil
}

func main(){
    circuitBreaker := Breaker(SimulatedCircuit,2)
	for i:=0; i<5;i++{
		response,err := circuitBreaker(context.Background())
		if err!=nil{
			fmt.Printf("Attempt %d: Error - %s\n", i+1, err.Error())
		} else{
			fmt.Printf("Attempt %d: Success - %s\n", i+1, response)
		}
	}
	time.Sleep(time.Second)

}