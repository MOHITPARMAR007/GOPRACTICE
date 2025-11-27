package main

import "fmt"

type vertex struct {
	x int
	y int
}

func main() {

	v := vertex{5, 6}
	//	fmt.Println(v)
	//	fmt.Println(v.x)
	p := &v
	p.x = 55
	//fmt.Print(v)

	primes := [6]int{1, 3, 5, 7, 9, 11}
	var s []int = primes[1:4]
	fmt.Println(s)
}
