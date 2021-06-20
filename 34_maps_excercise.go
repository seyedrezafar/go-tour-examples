package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	fmt.Println(words)
	var result = make(map[string]int)
	for i := range words {
	   value, ok := result[words[i]]
	   if ok{
		result[words[i]] = value + 1
	   }else{
		result[words[i]] = 1
	   }
	}
	return result
}

func main() {
	wc.Test(WordCount)
}

