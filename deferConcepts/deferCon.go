package main

import (
	"fmt"
)

/*
defer is an keyWord in go which is used to push fucntion in stack when we write defer in front of anything is
it pushesh that thing in the stack and thus stack in an lifo opration data structure it works on that
the last fucntion will run frist and the first function will execute in last
*/
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("DONE")
}
