package main

import "fmt"

/*
go does't have while or do while loops it has only for loops in that according to that we can modify
for loops in our own way below is the example of basic for loops
*/
func main() {
	/*	sum := 0
		for i := 1; i < 10; i++ {
			fmt.Println(sum)
			sum += i
		}
	*/
	// this is go's while loop
	sum2 := 0
	for sum2 < 10 {
		sum2 += sum2 + 1
	}
	fmt.Println(sum2)
}
