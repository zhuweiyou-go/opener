# opener

```bash
go get -u github.com/zhuweiyou-go/opener
```

```go
package main

import (
	"github.com/zhuweiyou-go/opener"
)

func main() {
	// open url
	// ignore error
	_ = opener.Open("https://github.com")

	// open directory
	err := opener.Open("c://")
	if err != nil {
		// handler error
	}
}
```
