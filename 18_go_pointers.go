package main

import "fmt"


func main(){
	i := 10

	p := &i
	fmt.Println("p is", *p)
	*p = 21
	fmt.Println("i after changing *p is: ", i)

}
