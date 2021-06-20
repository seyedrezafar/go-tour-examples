package main

import "fmt"


func main(){
	a := make([]int, 0, 10) // create an slice with length =0 and capacity = 10
	// a[5] = 30 // this will produce an error
	fmt.Println(a)
}
