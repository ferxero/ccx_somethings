package main

import (
	"fmt"

	"github.com/ferxero/ccx_somethings/count"
)

func main() {
	fmt.Println("", count.Max(1, 2))
	fmt.Println("", count.MaxOf3(1, 2, 3))
	fmt.Println("", count.MaxOf3(1, 3, 2))

	fmt.Println("", count.MaxOf3(2, 1, 3))
	fmt.Println("", count.MaxOf3(2, 3, 1))

	fmt.Println("", count.MaxOf3(3, 1, 2))
	fmt.Println("", count.MaxOf3(3, 2, 1))

	fmt.Println("", count.MaxOfN([]int{1, 4, 3, 2}))

}
