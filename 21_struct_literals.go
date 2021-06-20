package main

import "fmt"

type Vertex struct {
	X int
	Y int
}


func main(){
	v := Vertex{1, 2}
	// define a Vertex with just x initialization, the value of y would be 0
	v2 := Vertex{X: 1}
	fmt.Println("The value of v2.Y is ", v2.Y)
	p := &v
	fmt.Println("The value of p.x is: ", p.X)
	fmt.Println("The value of (*p).x is : ", (*p).X)
}
