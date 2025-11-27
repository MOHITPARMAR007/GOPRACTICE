package main

import (
	"fmt"
)

func main() {
	fmt.Println("slices ")
	slice := []int{2, 3, 4, 5, 6}
	//fmt.Println(slice[3])
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	//fmt.Println(slice)

	for i := len(slice) - 1; i >= 0; i-- {
		//		fmt.Println(slice[i], " ")
	}
	//fmt.Println(slice[1:4]) // 1 is inclusive and 4 is exclusive
	slice = append(slice[:2], slice[3:]...)
	//	fmt.Println(slice)

	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	slice2[0] = 0
	//	fmt.Println(slice2)

	makeSliceMax := []int{6, 2, 4, 1, 9}
	fmt.Println(makeSliceMax)
	var max int = 0
	i := 0
	for i < len(makeSliceMax) {
		if makeSliceMax[i] > max {
			max = makeSliceMax[i]
		}
		i++
	}
	fmt.Println(max)
	s := []int{}

	for len(s) < 7 {
		s = append(s, 1)
		fmt.Println("Len:", len(s), "Cap:", cap(s))
	}
}
