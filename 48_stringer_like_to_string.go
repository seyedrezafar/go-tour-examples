package main


import "fmt"


type Person struct {
  Name string
  Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}


func main(){
	p := Person{"Mohammad", 30}
	fmt.Println(p)
}
