package main
import (
	"fmt"
	"time"
	"sync"
	"context"
)

type Effector func (ctx context.Context)(string,error)

func Throttle(e Effector, max uint, refill uint, d time.Duration) Effector {
	var tokens = max
	var once sync.Once
	return func(ctx context.Context) (string, error) {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}
		once.Do(func() {
			ticker := time.NewTicker(d)
			go func() {
				defer ticker.Stop()
				for {
					select {
					case <-ctx.Done():
						return
					case <-ticker.C:
						t := tokens + refill
						if t > max {
							t = max
						}
						tokens = t
					}
				}
			}()
		})
		if tokens <= 0 {
			return "", fmt.Errorf("too many calls")
		}
		tokens--
		return e(ctx)
	}
}

func SimulateEffector(ctx context.Context)(string,error){
	time.Sleep(time.Millisecond*100)
	return "Done",nil

}

func main(){
	throttledEffector := Throttle(SimulateEffector,3,1,500*time.Millisecond)
	ctx := context.Background()
	var wg sync.WaitGroup 

	for i:=0;i<10;i++{
		wg.Add(1)
		go func(id int){
			defer wg.Done()
			result,err := throttledEffector(ctx)
			if err != nil {
				fmt.Printf("Error for worker %d: %v\n",id,err)
				return
			}
			fmt.Printf("Worker %d: %s",id,result)
		}(i)
	}
	wg.Wait()
}