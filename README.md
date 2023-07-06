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

### 017. 404页面(Not Found Page)

设置 404 页面有两种方式

第一种是设置响应头状态码为404，显示页面信息。
```go
w.WriteHeader(http.StatusNotFound)
fmt.Fprint(w, "Page not found")
```

第二种是同时设置状态码和页面内容。
```go
http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
```

### 018. http.Handler类型(The http.Handler Type)

定义Router并实现ServeHTTP
```go
type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
```

注册Router
```go
func main() {
	var router Router
	http.ListenAndServe(":3000", router)
}
```

### 019. http.HandlerFunc类型(The http.HandlerFunc Type)

将函数转换为 `http.HandlerFunc`
```go
func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
```

### 020. 探索处理转换(Exploring Handler Conversions)

`http.HandleFunc()` 底层本质也是调用的 `http.HandlerFunc()` 进行函数转换的，所以两种方式实现的效果是相同的。

### 021. FAQ练习(FAQ Exercise)

- 添加 `faqHandler` 处理函数。
- 添加 `/faq` 路径判断并调用处理函数。

## 三、路由和第三方库(Routers and 3rd Party Libraries)

### 022. 定义我们的路由需求(Defining our Routing Needs)

```
GET /galleries
POST /galleries

GET /galleries/:id
GET /galleries/:id/edit
DELETE /galleries/:id
```

### 023. 使用Git(Using git)

下载并安装[Git](https://git-scm.com/downloads)

创建分支

```bash
git checkout -b using-git
```

将所有文件的改动提交至暂存区

```bash
git add -A
```

提交暂存区的内容到本地仓库中

```bash
git commit -m "添加提交内容"
```

切换分支
```bash
git checkout casts
```

将 `using-git` 代码合并到 `casts` 分支中

```bash
git merge using-git
```

查看git状态
```bash
git status
```

显示暂存区的内容和被修改但未提交到暂存区文件的区别

```bash
git diff
```

推送本地仓库到远程仓库中

```bash
git push
```

### 024. 安装Chi(Installing Chi)

安装 `chi`
```bash
go get -u github.com/go-chi/chi/v5
```

清理未使用的依赖性
```bash
go mod tidy
```

### 025. 使用Chi(Using Chi)


导入`go-chi`包
```go
import (
	"github.com/go-chi/chi/v5"
)
```

创建路由
```go
func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000 ...")
	http.ListenAndServe(":3000", r)
}
```

### 026. Chi练习(Chi Exercises)

练习-添加URL参数

阅读文档，看看您是否可以添加一个URL参数到您的路由之一。在处理程序中检索它，并将其输出到结果HTML。
提示: 请参阅[这些文档](https://github.com/go-chi/chi#url-parameters) ，如果你需要一些指导。您不需要使用上下文，只需要使用URLParam方法

练习-使用内置中间件进行实验
Chi提供了不少内置中间件。一个是Logger中间件它将跟踪每个请求所花费的时间。尝试将其添加到您的应用程序中，然后仅添加到单个路由。

## 四、模板(Templates)

### 027. 什么是模板(What are Templates?)

模板是一种把获取某种内容的方式(例如html)并填充动态数据。

### 028. 我们为什么要使用服务端渲染(Why Do We Use Server Side Rendering?)

#### 服务端渲染

服务端渲染是由服务端返回html内容。

定义模板
```html
<body>
	<a href="/account">{{.Email}}</a>
	<h1>Hello, {{.Name}}!</h1>
</body>
```

服务端返回
```html
<body>
	<a href="/account">grayjunzi@email.com</a>
	<h1>Hello, grayjunzi!</h1>
</body>
```

#### 客户端渲染

客户端渲染是由客户端拼接html内容。 

服务端返回json数据。
```json
{
	"name":"grayjunzi",
	"email":"grayjunzi@email.com"
}
```

前端使用json数据并生成html。
```js
import React fron 'react';

function Example() {
	// 从服务端获取数据
	const {name, email} = fetchData();

	return (
		<body>
			<h1>Hello, {name}!</h1>
			<a href="/account">{email}</a>
		</body>
	);
}
```

### 029. 创建我们的第一个模板(Create Our First Template)

创建 `gohtml` 模板文件，并标记数据位置。
```html
<H1>Hello, {{.Name}}</H1>
```

解析 `gohtml` 模板文件，并填充数据。
```go
type User struct {
	Name string
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Smith",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
```

### 030. 跨站脚本攻击(Cross Site Scripting, XSS)

跨站脚本攻击指将恶意代码注入到网站内容当中。

导入 `html/template` 包时，模板会自动将html中的数据内容部分进行编码，防止跨站脚本攻击。
```go
import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Smith",
		Bio:  `<script>alert("Haha, you haven been h4x0r3d!")</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
```

### 031. 备选模板库(Alternative Template Libraries)

- `plush` - 语法与 ruby on rails 中的模板相似。

### 032. 上下文编码(Contextual Encoding)

go模板会自动根据上下文内容进行相应编码，如下在js脚本中字符串内容会自动添加双引号，而整型则不会添加双引号。
```html
<H1>Hello, {{.Name}}</H1>
<p>Bio: {{.Bio}}</p>

<script>
    const user = {
        "name": {{.Name}},
        "bio":  {{.Bio}},
        "age":  {{.Age}},
    }
    console.log(user);
</script>
```

生成的html内容如下
```html
<H1>Hello, John Smith</H1>
<p>Bio: &lt;script&gt;alert(&#34;Haha, you haven been h4x0r3d!&#34;)&lt;/script&gt;</p>
<p>Bio: 123</p>

<script>
    const user = {
        "name": "John Smith",
        "bio":  "\u003cscript\u003ealert(\"Haha, you haven been h4x0r3d!\")\u003c/script\u003e",
        "age":   123 ,
    }
    console.log(user);
</script>
```

### 033. 主页模板(Home Page via Template)

使用 `filepath` 拼接文件路径，并处理模板解析失败的情况。
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("解析模板: %v", err)
		http.Error(w, "解析模板出错.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("执行模板: %v", err)
		http.Error(w, "解析模板出错.", http.StatusInternalServerError)
		return
	}
}
```

### 034. 联系页面模板(Contact Page via Template)

编写代码时请考虑DRY原则(Don't Repeat Yourself)，编写相同的代码需要创建一个函数来执行该代码。

将通用部分抽离出来放到一个单独的函数中。
```go
func executeTemplate(w http.ResponseWriter, tplPath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("解析模板: %v", err)
		http.Error(w, "解析模板出错.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("执行模板: %v", err)
		http.Error(w, "解析模板出错.", http.StatusInternalServerError)
		return
	}
}
```

调用及修改时将会变得很容易。
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}
```

### 035. 常见问题页面模板(FAQ Page via Template)

可以将拼接路径的代码直接放到参数中执行。
```go
func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}
```