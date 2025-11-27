package main

import (
	"fmt"
)

func main() {
	score := map[string]int{
		"mohit":  23,
		"mahima": 23,
	}
	// can't add value in nil map and map is and referce type data structure
	m := make(map[string]int)
	m["a"] = 1
	fmt.Println(m)
	fmt.Println(score)
}
