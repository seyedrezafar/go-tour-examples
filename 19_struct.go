package main

import "fmt"

type Vertex struct {
	X int
	Y int
}


func main(){
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println("The value of X is : ", v.X)
}
