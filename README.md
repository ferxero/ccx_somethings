# Library: ccx somethings
It's to do any something in go with modules.

## How use it

```zsh
$ mkdir some_app && cd some_app
$ go mod init gitccx.com/user_x/appx # It is a sample
$ touch main.go
```

## Create a file called main.go

```go
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
```

## Run and import automatically the library

```zsh
$ go run .
```

## Then you can run the test

```zsh
go test -v ./count
```