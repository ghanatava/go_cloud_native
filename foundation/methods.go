package main 
import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Square(){
	v.X *=  v.X
	v.Y *= v.Y

}

type myMap map[string]int 

func (m myMap) Length() int{
	return len(m)
}

func main(){
	vert := &Vertex{3,1}
	fmt.Println(vert)

	vert.Square()
	fmt.Println(vert)

	mm := myMap{"hi": 0, "you gopher yet":1}
	fmt.Println(mm)
	fmt.Println(mm["you gopher yet"])
	fmt.Println(mm.Length())

}