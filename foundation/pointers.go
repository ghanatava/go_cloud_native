package main
import "fmt"

func main(){

    var a int = 5
	var p *int = &a

	fmt.Println(*p)
	fmt.Println(p)

	*p=20
	fmt.Println(*p)
	fmt.Println(a)
    
	var n *int
	var x, y int
	fmt.Println(n) // "<nil>"
	fmt.Println(n == nil) // "true" (n is nil)
	fmt.Println(x == y) // "true" (x and y are both zero)
	fmt.Println(&x == &x) // "true" (*x is equal to itself)
	fmt.Println(&x == &y) // "false" (different vars)
	fmt.Println(&x == nil) // "false" (*x is not nil)

}