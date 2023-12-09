package main
import (
	"time"
	"context"
	"errors"
	"sync"
	"fmt"
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

func DebounceFirst(circuit Circuit,d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(ctx context.Context)(string,error){
		m.Lock()
		defer func(){
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		if time.Now().Before(threshold){
			return result,err
		}
		result,err = circuit(ctx)
		return result,err
	}
}

func SimulatedCircuit(ctx context.Context)(string,error){
	attempts:=2
	if attempts > 0 {
		attempts--
		return "",errors.New("Simulated Response")
	} else{
		return "Success",nil
	}
}


func main(){	
	wrapped := Breaker(DebounceFirst(SimulatedCircuit, 500*time.Millisecond), 2)
	for i := 1;i<5;i++{
		start := time.Now()
		result,err := wrapped(context.Background())
		elapsed := time.Since(start)

		if err!=nil{
			fmt.Printf("Error %v\n",err)
		}else{
			fmt.Printf("Call %d: Result: %s ElapsedTime: %s\n",i,result,elapsed)
		}
	}

	time.Sleep(time.Millisecond*200)
}