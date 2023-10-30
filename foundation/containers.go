package main
import "fmt"

func main(){
	var a [3]int // Zero-value array of type [3]int
	fmt.Println(a) // "[0 0 0]"
	fmt.Println(a[1])
    
	a[1] = 42 // Update second index
	fmt.Println(a) // "[0 42 0]"
	fmt.Println(a[1]) // "42"
	i := a[1]
	fmt.Println(i)

    n := make([]int,3)
	fmt.Println(n) // "[0 0 0]"
	fmt.Println(len(n)) // "3"; len works for slices and arrays
	n[0] = 8
	n[1] = 16
	n[2] = 32
	fmt.Println(n)
	fmt.Println(n[:1])
	fmt.Println(a[1:])
    
}
