package main

import "fmt"


func Sqrt(x float64) float64 {
	z := float64(1)
	previous_z := float64(0)
	for  {
		z -= (z*z - x) / (2*z)
		if z == previous_z {
			break
		}else{
			previous_z = z
		}
		fmt.Println("The value of z in teration x is: ", z)
	}
	return z
}


func main() {
	fmt.Println(Sqrt(4))
}
