package main

import (
	"fmt"
)

func change(x int) {
	x = 20
}
func change2(x *int) {
	*x = 20
}
func main() {
	//	fmt.Println("hello_-world")
	/*i, j := 10, 20
	p := &i        // this is how we decakre pointer in the go
	fmt.Println(p) // yaha pr kuch bhi garbage value print karega yh
	*p = 22        // this will change the value of p because this is pointing to the address of the i
	fmt.Println(i, j)
	*/

	i := 10
	change(i)
	fmt.Println(i) // still the value will not change beacuse in change funtion there is and copy of i
	// is passed not the real i value so solution is pointer lets make another fun with refernce
	change2(&i)
	fmt.Println(i)
}
