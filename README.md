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

### 036. 模板练习(Template Exercises)

1. 模板中显示嵌套结构数据。
2. 模板中遍历Slice数据。
3. 模板中进行判断操作。

## 五、代码组织方式(Code Organization)

### 037. 代码组织方式(Code Organization)

我们的代码变得相当冗长，长期将所有代码都放到 `main.go` 文件中对于后期维护及查找问题都会造成很大的困扰。

一个好的代码结构应该是：
- 更容易找到问题。
- 更容易添加新功能。

#### 扁平结构(Flat Structure)

所有代码都在一个包中，用文件来分隔代码
```
myapp/
	gallery_store.go
	gallery_handler.go
	gallery_templates.go
	user_store.go
	user_handler.go
	user_templates.go
	router.go
	...
```

#### 关注点分离(Separation of concerns)

根据职责来划分代码。

Model-View-Controlle(MVC) 是采用这种策略的一种流行的结构。

- `models` - 数据、逻辑、规则，通常是数据库。
- `views` - 渲染一些东西，通常是html。
- `controller` - 把它连接起来。接受用户输入，将其传递给模型以完成操作，然后将数据传递给视图以呈现事物，通常是处理程序。

```
myapp/
	controllers/
		gallery_handler.go
		user_handler.go
		...
	views/
		gallery_templates.go
		user_templates.go
		...
	models/
		gallery_store.go
		userr_store.go
		...
```

#### 依赖型结构(Dependency-based structure)

基于依赖关系进行结构化，但是它们都使用一组通用的接口和类型。

```
myapp/
	user.go
	user_store.go

	psql/
		user_store.go
```

#### 其他结构

- 领域驱动设计(Domain-driven Design, DDD)
- 洋葱架构(Onion architecture)

### 038. MVC概述(MVC Overview)

`Model-View-Controller` - 又名MVC是一种代码架构模式。

根据职责组织代码，常被成为关注点(关注点分离)。

#### 模型(Models)

模型是负责关于数据、逻辑与规则的部分。

这通常意味着与数据库进行交互，但也可能意味着与来自其他服务或API的数据进行交互。通常包括验证数据、规范化数据等。

例如，我们的Web应用程序将有用户帐户，用于验证密码和验证用户身份的逻辑都将在models包中。


#### 视图(Views)

视图一般用于渲染html页面。

API可以使用MVC，视图可以负责生成JSON。

尽可能少的逻辑。只需要呈现数据的逻辑。

#### 控制器(Controllers)

控制器用来连接模型与视图的。

它不会直接渲染HTML，也不会直接触及数据库，但是它会调用models和views包中的代码来做这些事情。

控制器中不应该有太多的逻辑，而是将数据传递到应用程序的不同部分，这些部分实际上处理执行需要完成的任何工作。

### 039. 用MVC完成一个Web请求(Walking Through a Web Request with MVC)

1. 用户提交联系人信息更新请求。
2. Router转发至`UserController`中。
3. `UserController`使用`UserStore`更新用户联系人信息。
4. `UserStore`返回更新后的用户数据。
5. `UserController`使用`ShowUser`视图生成HTML响应。
6. `ShowUser`视图写入HTMl响应到 `http.ResponseWriter` 中。

### 040. MVC练习(MVC Exercises)

**1. MVC代表什么意思？**

MVC是`Model-View-Controller`的缩写。

**2. MVC各层职责是什么？**

- `Model` - 负责数据、逻辑与规则，一般是与数据库交互。
- `View` - 负责渲染页面，一般是渲染html。
- `Controller` - 负责连接模型与视图，调用`models`处理业务逻辑并将返回数据交给`views`进行渲染。

**3. 使用MVC的好处和坏处是什么？**

**4. 阅读其他构建代码的方法**

## 六、开始应用MVC(Starting to Apply MVC)

### 041. 创建Views包(Creating the Views Package)

将解析与渲染模板拆分成两个独立的函数。
```go
type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("执行模板: %v", err)
		http.Error(w, "解析模板出错.", http.StatusInternalServerError)
		return
	}
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTpl: tpl,
	}, nil
}
```

### 042. fmt.Errorf

创建一个新错误
```go
errors.New("connection failed")
```

格式化错误信息
```go
fmt.Errorf("failed: %w", err)
```

### 043. 启动时验证模板(Validating Template at Startup)

创建模板处理函数
```go
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
```

为每个路由设置这个处理函数，将会在程序启动时自动验证模板是否存在问题。
```go
tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
if err != nil {
	panic(err)
}
r.Get("/", controllers.StaticHandler(tpl))
```

### 044. Must函数(Must Functions)

创建Must函数，在发生错误时自动进行panic。
```go
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
```

使用方式
```go
r.Get("/", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))
```

### 045. 练习(Exercises)

1. 使用新模式向您的应用添加新的静态页面。
2. 试验一下错误。

```go
errNotFound := errors.New("not found")

errors.Is(err, errNotFound)

errors.As(err, &target)
```

## 七、提高我们的视图(Enhancing our Views)

### 046. 嵌入模板文件(Embedding Templates Files)

创建`fs.go`文件，通过 `//go:embed` 指令将静态资源文件打包至编译后的文件中。
```go
package templates

import "embed"

//go:embed *
var FS embed.FS
```

添加 `ParseFS` 函数，通过文件系统解析模板。
```go
func ParseFS(fs fs.FS, pattern string) (Template, error) {
	tpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTpl: tpl,
	}, nil
}
```

使用 `ParseFS` 加载模板。
```go
r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml"))))
```

### 047. 可变参数(Variadic Parameters)

使用 `...int` 接收可变参数，这个参数将是一个数组。
```go
func add(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
```

调用`add`函数时，参数的数量是可变的。
```go
fmt.Println( add(1, 2, 3, 4) )
```

### 048. 命名模板(Named Templates)

命名模板就像组件，可以多次重用它们，可以把多种东西放到页面上，只需要写一次HTML。

定义命名模板
```html
{{define "lorem"}}
<p>
	Lorem ipsum, dolor sit amet consectetur adipisicing elit. Quam quae ea nobis reiciendis ratione eligendi quis
    perferendis dignissimos consequatur, in sapiente? Omnis adipisci, soluta tempora aliquam eos earum et magnam.
</p>
{{end}}
```

使用命名模板
```html
{{template "lorem"}}
```

### 049. 动态FAQ页面(Dynamic FAQ Page)

添加`FAQ`处理函数，并将数据传入模板中。
```go
func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free trial for 30 days on any paid plans.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have support staff answering emails 24/7, though response times may be abit slower on weekends.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `<a href="mailto:supportalenslocked.com">Email us</a>`,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
```

在路由上应用`FAQ`处理函数。
```go
r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))
```

定义模板并遍历数据
```html
<h1>FAQ Page</h1>
<ul>
    {{range .}}
        {{template "qa" .}}
    {{end}}
</ul>

{{define "qa"}}
<li><b>{{.Question}}</b>{{.Answer}}</li>
{{end}}
```

### 050. 可复用布局(Reusable Layouts)

有两种方式，第一种是定义布局模板。
```html
{{define "header"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Lenslocked</title>
</head>

<body>

{{end}}

{{define "footer"}}

</body>

</html>
{{end}}
```

然后在页面中引用局部模板。
```html
{{template "header"}}

<h1>欢迎来到我的站点!</h1>

{{template "footer"}}
```

最后将模板加入到路由中。
```go
r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))))
```

第二种方式是在布局模板也中指定加载的模板。
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Lenslocked</title>
</head>

<body>
    {{template "page" .}}
</body>

</html>
```

然后在页面中定义模板。
```html
{{define "page"}}
<h1>欢迎来到我的站点!</h1>
{{end}}
```

最后将模板加入到路由中，注意顺序，要先指定布局模板，再指定具体页面的模板。
```go
r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "layout-parts.gohtml", "home.gohtml"))))
```

### 051. Tailwind CSS

添加模板页，以CDN的方式引入`TailwindCSS`。
```html
{{define "header"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tailwind</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body>
{{end}}

{{define "footer"}}
</body>

</html>
{{end}}
```

### 052. 实用性优先的CSS(Utility-first CSS)

在模板页面中添加TailwindCSS样式。

### 053. 添加导航栏(Adding a Navigation Bar)

添加导航栏样式
```html
<header class="bg-gradient-to-r from-blue-800 to-indigo-800 text-white">
	<nav class="px-8 py-6 flex items-center">
		<div class="text-4xl pr-8 font-serif">Lenslocked</div>
		<div class="flex-grow">
			<a class="text-lg font-semibold hover:text-blue-100 pr-8" href="/">Home</a>
			<a class="text-lg font-semibold hover:text-blue-100 pr-8" href="/contact">Contact</a>
			<a class="text-lg font-semibold hover:text-blue-100 pr-8" href="/faq">FAQ</a>
		</div>
		<div>
			<a class="pr-4" href="#">Links</a>
			<a class="px-4 py-2 bg-blue-700 hover:bg-blue-600 rounded" href="#">Sign in / Sign up</a>
		</div>
	</nav>
</header>
```

### 054. 练习(Exercises)

1. 使用TailwindCSS，并尝试设计一些东西。
2. 创建某种类型的 `<footer>` 模板，并尝试确保它包含在所有视图中。
3. 请查看嵌入式软件包，看看您还可以用它做些什么。

## 八、注册页面(The Signup Page)

### 055. 创建注册页面(Creating the Signup Page)

添加 `signup.gohtml` 注册模板页。
```html
{{template "header" .}}

<form action="/users" method="post">
    <div>
        <label for="email">邮箱</label>
        <input name="email" id="email" type="email" placeholder="邮箱地址" required autocomplete="email" />
    </div>
    <div>
        <label for="password">密码</label>
        <input name="password" id="password" type="password" placeholder="密码" required />
    </div>
    <div>
        <button>注册</button>
    </div>
</form>

{{template "footer" .}}
```

### 056. 设置注册页面的样式(Styling the Signup Page)

为注册表单添加样式
```html
<div class="py-12 flex justify-center">
    <div class="px-8 py-8 bg-white rounded shadow">
        <h1 class="pt-4 pb-8 text-center text-3xl font-bold text-grray-900">
            今天开始分享你的照片！
        </h1>
        <form action="/users" method="post">
            <div class="py-2">
                <label for="email" class="text-sm font-semibold text-gray-800">邮箱</label>
                <input name="email" id="email" type="email" placeholder="邮箱地址" required autocomplete="email"
                    class="w-full px-3 py-2 border borderr-gray-300 placeholder-gray-500 text-gray-800 rounded" />
            </div>
            <div class="py-2">
                <label for="password" class="text-sm font-semibold text-gray-800">密码</label>
                <input name="password" id="password" type="password" placeholder="密码" required
                    class="w-full px-3 py-2 border borderr-gray-300 placeholder-gray-500 text-gray-800 rounded" />
            </div>
            <div class="py-4">
                <button
                    class="w-full py-4 px-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded font-bold text-lg">注册</button>
            </div>
            <div class="py-2 w-full flex justify-between">
                <p class="text-xs text-gray-500">已经有帐号了？<a href="/signin" class="underline">登录</a></p>
                <p class="text-xs text-gray-500"><a href="/reset-passworrd" class="underline">忘记密码？</a></p>
            </div>
        </form>
    </div>
</div>
```

### 057. REST简介(Intro to REST)

REST 是 `REpresentational State Transfer` 的缩写，它是一种架构风格(architectural style)。

- 无状态(Stateless)，我们不需要记住客户端在做什么。一个请求有足够的信息供我们响应。
- 客户端和服务器交互，REST原则可以帮助我们设计一个更直观的客户端使用的服务器。

请求包括：
- 请求方法(HTTP Method) 也称作请求动词(HTTP verb)。
- 请求路径。

REST端点(endpoints)与资源有关
| **HTTP Method** | **路径** | **作用** |
| `GET` | /galleries | 读取相册列表 |
| `GET` | /galleries/:id | 读取单个相册 |
| `POST` | /galleries | 创建一个相册 |
| `PUT` | /galleries/:id | 修改一个相册 |
| `DELETE` | /galleries/:id | 删除一个相册 |

### 058. 用户控制器(Users Controller)

创建Users控制器
```go
package controllers

import (
	"net/http"

	"github.com/grayjunzi/lenslocked/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
```

初始化Users控制器并指定注册页面模板。
```go
usersC := controllers.Users{}
usersC.Templates.New = views.Must(views.ParseFS(
	templates.FS,
	"signup.gohtml", "tailwind.gohtml",
))

r.Get("/signup", usersC.New)
```

### 059. 使用接口解耦(Decouple with Interfaces)

抽象模板接口
```go
package controllers

import "net/http"

type Template interface {
	Execute(w http.ResponseWriter, data interface{})
}
```

### 060. 解析注册表单(Parsing the Signup Form)

调用 `ParseForm()` 函数解析表单，使用 `PostForm.Get()` 或 `FormValue()` 获取表单值。
```go
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Email: %s, Password: %s\n", r.PostForm.Get("email"), r.PostForm.Get("password"))
	fmt.Fprintf(w, "Email: %s, Password: %s\n", r.FormValue("email"), r.FormValue("password"))
}
```

### 061. URL查询参数(URL Query Params)

可通过 `FormValue()` 函数获取 `query string` 中的值。
```go
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}
```

模板判断值是否存在并自动聚焦。
```html
<div class="py-2">
	<label for="email" class="text-sm font-semibold text-gray-800">邮箱</label>
	<input name="email" id="email" type="email" placeholder="邮箱地址" required autocomplete="email"
		value="{{.Email}}" {{if not .Email}}autofocus{{end}}
		class="w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-800 rounded" />
</div>
<div class="py-2">
	<label for="password" class="text-sm font-semibold text-gray-800">密码</label>
	<input name="password" id="password" type="password" placeholder="密码" required {{if
		.Email}}autofocus{{end}}
		class="w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-800 rounded" />
</div>
```

### 062. 练习(Exercises)

1. 在表单中添加新的数据，并向服务器提交数据，尝试在服务器上解析它们。
2. 调整路径，例如，将注册页面改为 `/users/new`，将注册提交移动到 `/signup`，查看需要更改哪些内容才能实现此操作。

## 九、数据库与PostgreSQL

### 063. 数据库简介(Intro to Databases)

我们需要一些存储数据的方法来使应用程序成为一个真正的Web应用程序，而不单单只是静态页面。

数据库的类型很多，都是针对不同的情况而设计的:
- 关系数据库(Relational Databases) - `PostgreSQL` 和 `MySQL`
- 文档存储库(Document Stores) - `MongoDB`
- 图形数据库(Graph Databases) - `Dgraph` 和 `Neo4i`
- 键值对存储(key/value stores) - `BoltDB` 和 `etcd`

每个DB都有优点(pros)和缺点(cons)。如果一个数据库在一件事上做得很好，那么它就是在某个地方做了一个权衡。因此，大多数大公司使用各种数据库来完成不同的任务。

我们将使用PostgreSQL，它是:
- 非常受欢迎。
- 免费和开放源代码。
- 规模(Scales)非常好。
- 不太复杂。
- 在`Go`中得到了很好的支持。
- 防止竞争条件的事务。

关系型数据库被广泛使用。
- 通过Docker运行Postgres
- 与Postgres交互
- 了解Postgres的工作原理

### 064. 安装Postgres(Installing Postgres)

查看docker版本
```bash
docker version
```

查看docker-compose版本
```bash
docker-compose version
```

定义`docker-compose.yml`文件，运行`postgres`与`adminer`两个服务。
```yml
version: '3.9'

services:

  db:
    image: postgres
    restart: always
    environment:
      - POSTGRES_USER=root
      - POSTGRES_PASSWORD=root
      - POSTGRES_DB=lenslocked
    ports:
      - 5432:5432

  adminer:
    image: adminer
    restart: always
    environment:
      - ADMINER_DESIGN=dracula
    ports:
      - 3333:8080
```

执行以下命令，将会根据 `docker-compose.yml` 文件中的配置，拉取并启动两个镜像。
```bash
docker-compose up
```

> 执行该命令时需与 `docker-compose.yml` 在同一目录下。

访问 `http://localhost:3333` 打开 `adminer` 数据库管理界面。

在命令中增加 `-d` 可在后台运行容器。
```bash
docker-compose up -d
```

使用以下命令移除 `docker-compose.yml` 文件中配置的所有镜像。
```bash
docker-compose down
```

### 065. 连接到Postgres(Connecting to Postgres)

执行以下命令进入容器，其中 `-it` 指定要进入的容器，`-U` 指定数据库的用户名，`-d` 指定数据库。
```bash
docker exec -it lenslocked-db-1 /usr/bin/psql -U root -d lenslocked
```

> `-i` 全称为 `--interactive` 用于以交互模式运行容器，`-t` 全称为 `--tty` 会为容器重新分配一个伪输入终端。

### 066. 创建SQL表(Creating SQL Tables)

执行以下sql语句创建`users`表。
```sql
CREATE TABLE users (
	id SERIAL,
	age INT,
	name TEXT,
	email TEXT
);
```

### 067. Postgres数据类型(Postgres Data Types)

`Types`提供了一种方法来定义要在列中存储的数据类型。

[查看PostgreSQL中提供的数据类型](https://postgresql.org/docs/current/datatype.html)

| 类型 | 描述 |
| -- | -- |
| `int` | 这用于存储-2147483648和2147483647之间的整数。 |
| `serial` | 这是用来存储1和2147483647之间的整数。`int` 和 `serial`最大的区别在于，如果你没有提供一个值，serial会自动设置一个值，新值总是加1。这对于id列很有用，在id列中，您希望每一行都有一个唯一的值，并允许数据库决定使用什么值。 |
| `varchar` | 这类似于Go或其他编程语言中的字符串，只是我们必须告诉数据库我们存储的任何字符串的最大长度是多少。 |
| `text` | 这是一个特定于PostgreSQL的类型。并且可能不是在所有形式的SQL中都可用，但它在底层与varchar基本相同，但是在声明字段时不必指定最大字符串长度。 |
| `uuid` | 唯一Id |

### 068. Postgres约束(Postgres Constraints)

约束是我们可以应用于表中字段的规则。例如，我们可能希望确保数据库中的每个用户都有唯一的`id`，因此可以使用UNIQUE约束。

[查看PostgreSQL中提供的约束](https://postgresql.org/docs/current/static/ddl-constraints.html)

| 约束 | 描述 |
| -- | -- |
| `UNIQUE` | 这确保了数据库中字段的每个记录值都被设置为唯一的。 |
| `NOT NULL` | 这确保了数据库中的每条记录都有这个feed的值。当你不为一个字段提供值时，数据库通常会存储null，但这阻止了它的有效性。 |
| `PRIMARY KEY` | 这个约束类似于UNIOUE和NOT NULL的组合，但它只能在每个表上使用一次，并且它会自动导致为该字段创建索引。这个索引是用来使查找这个字段的记录更快捷。 |

### 069. 创建用户表(Creating a Users Table)

如果`users`表存在则删除该表。
```sql
DROP TABLE IF EXISTS users;
```

创建`users`表
```sql
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	age INT,
	name TEXT,
	email TEXT UNIQUE NOT NULL
);
```

### 070. 插入记录(Inserting Records)

插入表数据
```sql
INSERT INTO users VALUES(1, 25, 'grayjunzi', 'grayjunzi@email.com');
```

插入数据指定字段
```sql
INSERT INTO users (age, name, email) VALUES(30, 'admin', 'admin@email.com');
```

查询用户表
```sql
SELECT * FROM users;
```

### 071. 查询数据(Querying Records)

查询表中所有数据
```sql
SELECT * FROM users;
```

查询表中所有数据并指定返回列
```sql
SELECT id, email FROM users;
```

### 072. 过滤数据(Filtering Queries)

根据条件筛选数据
```sql
SELECT * FROM users WHERE email='admin@email.com';
SELECT * FROM users WHERE age > 20;
```

根据`age`和`name`条件筛选并返回1条数据。
```sql
SELECT * FROM users WHERE age > 25 OR name = 'grayjunzi' LIMIT 1;
```

### 073. 更新数据(Updating Records)

```sql
UPDATE users SET age = 26 WHERE name = 'grayjunzi';
```

### 074. 删除数据(Deleting Records)

```sql
DELETE FROM users WHERE name = 'admin';
```

### 075. 其他SQL资源(Additional SQL Resources)

**Select Star SQL** - [https://selectstarsql.com](https://selectstarsql.com)

**SQL Murder Mystery** - [https://mystery.knightlab.com](https://mystery.knightlab.com)

**Codecademy's Learn SQL course** - [https://www.codecademy.com/lrn/learn-sql](https://www.codecademy.com/lrn/learn-sql)

**w3schools** - [https://w3schools.com](https://w3schools.com)

**Khan Academy's SQL course** - [https://www.khanacademy.org/computing/computer-programming/sql](https://www.khanacademy.org/computing/computer-programming/sql)

## 十、在Go中使用Postgres(Using Postgres with Go)

### 076. 使用Go连接到Postgres(Connecting to Postgres with Go)

安装 `pgx`
```go
go get github.com/jackc/pgx/v4
```

使用 `pgx` 连接Postgres数据库
```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=root password=root dbname=lenslocked sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
}
```

### 077. 导入的副作用(Imports with Side Effects)

Go中导入语句如果没有使用将会自动删除，此时需要加入`_`防止误删。
```go
import (
	_ "github.com/jackc/pgx/v4/stdlib"
)
```

> 导入的作用是为了执行该包下的 `init()` 函数，用于设置一个驱动程序，然后调用SQL寄存器