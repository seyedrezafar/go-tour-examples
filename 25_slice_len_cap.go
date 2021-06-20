package main


import "fmt"


func main(){
	slice := []int{1,2,3,4,5,6}
	fmt.Println(slice)
	subSlice := slice[1:3]
	fmt.Println(subSlice)
	fmt.Println("The capacity of sub slice is: ", cap(subSlice))
	fmt.Println("The length of sub slice is: ", len(subSlice))

}
