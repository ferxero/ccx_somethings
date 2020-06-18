package main

import (
	"fmt"

	"github.com/ferxero/ccx_somethings/count"
)

func main() {

	// get max value from 2 numbers
	fmt.Println("", count.Max(1, 2))

	// only used to 3 integer elements
	fmt.Println("", count.MaxOf3(1, 2, 3))
	fmt.Println("", count.MaxOf3(1, 3, 2))

	fmt.Println("", count.MaxOf3(2, 1, 3))
	fmt.Println("", count.MaxOf3(2, 3, 1))

	fmt.Println("", count.MaxOf3(3, 1, 2))
	fmt.Println("", count.MaxOf3(3, 2, 1))

	// now you can use it with n values
	fmt.Println("", count.MaxOfN([]int{1, 4, 3, 2}))

}
