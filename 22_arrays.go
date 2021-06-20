package main

import "fmt"


func main(){
	// the arrays could not be resized and could be created using [n]T, n indicate number of elements and T provide the type
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
