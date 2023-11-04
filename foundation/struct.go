package main

import "fmt"

type Vertex struct {
    x, y float64
}

func main() {
    var v Vertex
    fmt.Println(v) // Structs are never nil

    v = Vertex{1.0, 2.0}
    v.x *= 10
    v.y *= 5
    fmt.Println(v)

	var v1 *Vertex = &Vertex{3.2,1.4}
	fmt.Println(v1)
	v.x,v.y = v.y,v.x
	fmt.Println(v1)
}
