package main

import "fmt"

func counter() func() int {
	count := 0 // outer variable (local)

	return func() int { // inner function = closure
		count++ // remembers 'count'
		return count
	}
}

func main() {
	c := counter()

	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3
}
