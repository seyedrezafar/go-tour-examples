package main

import "fmt"


func main(){
	// slices are pointers to underlying array
	slice := []bool{true, true, false}
	fmt.Println(slice)
}

