package main


import "fmt"


type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell labs": Vertex{40.2, -74.3},
	"Google": Vertex{37.3, -122.22},
}

func main(){
	fmt.Println(m)
}
