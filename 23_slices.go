package main

import "fmt"


func main(){
	// the slices could be created using []T
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	x int = 3
	var primesSlices int[] = primes[1:4]
	fmt.Println(primesSlices)
}
