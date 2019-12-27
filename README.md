# hello-go
`hello go package` is very simple go package.  
Get string s and return s + `Hello World`.

## usage
`go get github.com/st34-satoshi/hello-go`

### main.go

```go
package main

import (
	"fmt"
	"github.com/st34-satoshi/hello-go"
)

func main() {
	fmt.Println(hello_go.WithComment("hello "))
}
```

then out put is  
`hello Hello World`

