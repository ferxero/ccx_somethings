# ccx_somethings
It's to do any something in go with modules.

# How use it

```sh
$ mkdir some_app && cd some_app
$ go mod init x.com/fx/appx
$ touch main.go
```

# Create a file called main.go

```sh
package main

import (
	"fmt"

	"github.com/ferxero/ccx_somethings/count"
)

func main() {
	fmt.Println("", count.Max(1, 2))
}
```

# Run and import automatically the library

$ go run .

