package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum + 2
	y = x + 1
	return
}

// now lets learn about varibale in go
// var sum , totalSum int this is how we can declare the varibale in go if all the varibale are
// of same type we can guve the type of the varibale in the end
// and if type are diffrent we have to provide the type of the variable like this
// var sum int , var total boolean
// in go a varibale is never in its un intiallized state or undefined state or null state  it will be only in its zero state
// example
var total int

// by default boolean value is false
var flag1, flag2 bool

// one more decalration form we have in go which is only valid in Inside a functoin
// i := 1; if we try to do that outside of the function it will give error

func main() {
	fmt.Println("My fav Number is :", rand.Intn(10))
	// this below line will not work because of we are using pi in an external package and
	// go allows us only capital latters variable is to accessable to other packages..!
	//fmt.Println(math.pi)
	fmt.Println(math.Pi) // this line will work as we are using pi and captial P
	fmt.Println(add(3, 4))
	fmt.Println(swap("hello", " Mohit"))
	a, b := swap("toji", "fushiguro")
	fmt.Println(a, b)
	var sum int = 5
	x, y := split(sum)
	fmt.Println(x, y)
	fmt.Println(total, flag1, flag2)
}
