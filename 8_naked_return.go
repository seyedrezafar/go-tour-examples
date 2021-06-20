package main

import "fmt"

func split(sum int) (x, y int) {
	// naked returns should not be used in longer functions 
	x = sum * 4 / 9
	y = sum - x
	return 
}


func main(){
	fmt.Println(split(17))
}

