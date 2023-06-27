# openner

```bash
go get -u github.com/zhuweiyou-go/openner
```

```go
package main

import (
	"github.com/zhuweiyou-go/openner"
)

func main() {
	// open url
	// ignore error
	_ = openner.Open("https://github.com")

	// open directory
	err := openner.Open("c://")
	if err != nil {
		// handler error
	}
}
```
