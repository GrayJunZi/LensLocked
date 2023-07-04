# Gomo

Web Development with Go
---
Learn to build real, production-grade web applications from scratch.

## 一、开始入门(Getting Started)

### 001. 一个基本的Web应用程序(A Basic Web Application)

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

### 003. 包与导入(Packages and Imports)

在Go语言中使用 `package` 来标识当前包名称。

```go
package main
```

> 在Go语言中有一个名为`main`的特殊的包，用于告诉go程序从从哪里开始执行。

使用`import` 关键字导入其他的包。

```go
import (
    "fmt"
    "net/http"
)
```

- `fmt` - 标准库，用于终端的输入输出及格式化等常用函数。
- `net/http` - 标准库，用来做任何与HTTP相关的事情。

### 004. 编辑器与自动导入(Editors and Automatic Imports)

VSCode 微软开发的免费且开源的编辑器，可通过安装插件的形式来方便Go语言开发可具有智能提示、自动导包、查看Go语言源码等众多功能。

### 005. 代码中的HelloWorld部分(The Hello, World Part of our Code)

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello, World")
}
```

### 006. Web请求(Web Request)

Web请求有几个主要组成部分：

- Url - 请求服务器地址。
- Mehtod - 请求方式。
- StatusCode - 状态码。
- RequestBody - 请求正文。
- ResponseBody - 响应正文。

### 007. HTTP方法(HTTP Methods)

有以下几种HTTP请求方法：

- GET - 获取一些资源。
- POST - 提交一些资源。
- PUT - 修改一些资源。
- DELETE - 删除一些资源。

### 008. 处理函数(Our Handler Function)

#### http.ResponseWriter 

它是一个接口，定义了一组方法，我们可以用来响应Web请求。我们可以用它来写响应体、Headers、Cookie、HTTP StatusCode 等。

#### http.Request

它是指向结构类型的指针，定义了传入请求的数据、请求正文、路径、请求方法、Headers等。

#### 为什么一个接口而一个是结构呢？

http.ResponseWriter是接口类型意味着可能传入不同的实现，例如使用`httptest`包可以易于测试代码。

### 009. 注册处理函数并启动Web服务(Registering our Handler Function and Starting the Web Server)

#### http.HandleFunc 

该方法用于注册处理函数。

#### http.ListenAndServe

该方法用于监听并启动Web服务。

### 010. Go模块(Go Module)

#### 依赖管理

确保其他开发人员可以使用类似的版本构建您的代码。

版本：
- 由三组数字组成，例如`v1.12.4`
- 主版本(Major)、次版本(minor), 补丁(patch)。


#### 在`GOPATH`路径之外运行

所有Go代码都位于计算机上的一个目录中（GOPATH）。
`Go Modules`允许我们在任何地方运行代码，只要我们初始化一个模块。

#### 设置Go Module

执行改命令将会创建`go.mod`文件。
```bash
go mod init github.com/grayjunzi/lenslocked
```

使用 `go get` 命令将会安装或更新依赖类库。

使用 `go mod tidy` 命令清理依赖。

## 二、添加新页面(Adding New Pages)

### 011. 动态重载(Dynamic Reloading)

安装 `modd`
```bash
go install github.com/cortesi/modd/cmd/modd@latest
```

创建 `modd.conf` 文件
```conf
**/*.go {
    prep: go test @dirmods
}

**/*.go !**/*_test.go {
    prep: go build -o ./ .
    daemon +sigterm: ./lenslocked.exe
}
```

执行命令
```bash
modd
```

### 012. 设置Header值(Setting Header Values)

调用 `w.Header().Set()` 设置响应头信息。
```go
func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
}
```

### 013. 创建Contact页面(Creating a Contact Page)

创建`contactHandler`函数。
```go
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at</p><a href='mailto:test@123.com'>test@123.com</a>")
}
```

注册路由
```go
http.HandleFunc("/contact", contactHandler)
```

### 014. 检查http.Request类型(Examing the http.Request Type)

在 `http.Request` 结构体中可以查看请求的URL路径。
```go
func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

func main() {
	http.HandleFunc("/", pathHandler)
	http.ListenAndServe(":3000", nil)
}
```

### 015. 自定义路由(Custom Routing)

判断请求路径跳转到不同处理函数中
```go
func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// TODO: 404 Not Found
	}
}
```

### 016. URL.Path与URL.RawPath对比(url.Path vs url.RawPath)

- `URL.Path` - 无论路径中是否有编码后的字符都会进行解码展示。
- `URL.RawPath` - 当路径中出现编码后的字符时才会有值，且会显示编码后的字符。


```go
func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
	fmt.Fprintln(w, r.URL.RawPath)
}
```