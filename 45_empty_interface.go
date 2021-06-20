package main

import "fmt"


func main(){
   var i interface{}
   describe(i)


   i = 42
   describe(i)

   i = "Hello world"
   describe(i)
}

func describe(i interface{}){
	fmt.Println("(%v, %T)", i , i)
}
