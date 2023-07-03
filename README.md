# Gomo

Web Development with Go
---
Learn to build real, production-grade web applications from scratch.

## 一、开始入门(Getting Started)

### 001. 一个基本的Web应用程序

```go
package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Starting the server on :3000 ...")
	http.ListenAndServe(":3000", nil)
}
```

### 002. 故障排除(Troubleshooting)