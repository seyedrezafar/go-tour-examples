package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0
	var fn = []int{0, 1}
	return func() int{
		res := 0
		if n <= 1 {
			res = fn[n]
		}else {
			res = fn[n-1] + fn[n-2]
			fn = append(fn, res)
		}
		n += 1
		return res

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}

