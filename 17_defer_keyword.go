package main

import "fmt"

func main(){
	// the execution of the defer statment will be postpone until the end of surrounding function 
	defer fmt.Println("World")

	defer fmt.Println("World 2")

	fmt.Println("Hello")
}
