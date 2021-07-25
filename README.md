# 背景

因为之前已经学习了 golang，所以急需一个项目来练练手，所以打算使用非常流行的 web 开发框架 gin + bootstrap 来开发一个简单博客系统

golang 在语法上与 c 非常相似，很多程序员的入门语言都是 c 语言，所以对于 c 比较熟悉的程序员在学习 golang 时候会比较得心应手，并且 golang 其中有一个显著的特点就是不允许指针运算，这大大提高了程序的安全程度，降低了编程难度，并且提供了可供 go 虚拟机自行分配调度的 goroutine 协程。在效率上能和 java 上一较高下，在内存占用上又比 java 厉害，同时在高并发场景表现显著，对于后端开发的上手难易程度来说也是非常人性化的。可惜 golang 的历史并不悠久，目前生态仅局限于在后端发光发亮，等待生态建立，未来还是很有机会和 java 争夺天下的

# 项目简介

**项目功能**

一个博客管理系统，用户可以注册和登录，也可以查看首页，查看所有博客，登录后的用户作为博客的后台管理者，可以写博客，编辑所有博客，删除所有博客，还可以查看关于我自己的信息

主页面调用了一些开放 api，所以首页展示会比较缓慢，我的页面存在我的个人信息，可以跳转我的 github，csdn，微信，邮箱等

# 使用环境

使用环境如下：

- 编程语言：go 1.16.2
- 包管理工具：go module
- Web 框架：gin
- 前端框架：bootstrap v3，jquery 3
- 编辑工具：intellij goland
- 数据库：mysql 5.7
- 分支管理：git 2.24.3
- 其他：前端语言 html/js，jquery 插件

# 项目结构

采用 web 开发中经典的 mvc 的思想，结构如下：

```
- config 「配置文件」
- controller 「控制器」
- database 「数据库底层操作」
- middleware 「gin 中间件」
- model 「对象层」
- router 「路由层」
- service 「对数据库操作的业务封装」
- static 「前端静态资源」
	- css 「css 文件」
	- img 「图片」
	- js 「js 文件」
- util 「工具类」
- view 「前端 html 界面」
- main.go 「主函数」
```

# main 主函数

主函数中主要进行三项操作，每次启动服务，通过运行 main 函数实现

- **初始化数据库**

  数据库初始化

- **创建路由实例**

  创建一个注册好的路由实例

- **运行服务**

  路由实例监听端口运行服务

# router 路由

在 java spring 框架中，路由是通过注解的形式放在控制器上方的，java 中的注解比较强大，运行了 application 时候相当于 go 中注册了路由，并将路由绑定到了控制器，并进行了监听，要知道 go 中路由的写明和绑定控制器需要程序员自己去写明

我的博客系统中所有的路由都是放在一个 go 的路由文件中，并且 package 是 router，其中主要做的 5 项操作

**创建路由实例**

要知道 gin 框架底层使用的就是 httprouter，所以通过 gin 框架的 gin 包可以产生一个默认的路由实例

```go
r := gin.Default()
```

**静态资源加载**

使用 Static 函数来加载前端的静态资源

**前端模版加载**

使用 LoadHTMLGlob 函数来家在前端 html 模版，之所以没有使用`view/*`的形式，是因为 view 包下面的 html 模版依据博客系统的模块进行了分类

```go
r.LoadHTMLGlob("view/**/*")
```

**中间件的使用**

在路由注册之前需要声明使用的中间件，这样在执行某个路由之前，就会先执行一下中间件的操作。这里使用的中间件是 session 的中间件，写法如下：

```go
// UseSession session 中间件
func UseSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte("loginUser"))
	return sessions.Sessions("loginUser", store)
}
```

中间件都需要返回`gin.HandlerFunc`，然后在 router 中通过如下方式完成中间件的声明

```go
r.Use(middleware.UseSession())
```

值得注意的是通过 GET 或者 POST 注册路由的函数中第二个参数的类型也是`gin.HandlerFunc`，注册路由时候第二个参数一般填写的是 controller 的函数名。中间件声明之后在 controller 中即可通过如下方式来保存用户 session 了

```go
// c *gin.Context，s Session
s := sessions.Default(c)
// 保存用户 session
s.Set("loginUser", username)
```

**路由注册**

接下来我们来做 GET/POST 路由注册，依据 restful 规范，我们划分了一下模块，主要注册了以下几大类的路由

- 首页路由
- 注册页路由
- 登录页路由
- 文章相关页面路由
- 关于我页面路由

注册路由中传递的参数是路由名称和路由需要绑定的 controller 函数名，值得注意的是 controller 函数名是属于`gin.HandlerFunc`类型。最后 InitRouter 函数返回`*gin.Context`类型路由，供 main 函数调用

# model 层

model 层主要存放了文章 model 和用户 model，包括其中的 struct 结构和操作各自 model 的基本数据库操作，比如说查询用户，依据 id 查询用户，查询文章总数，依据作者查询文章总数，分页查询文章总数......

**User**

```go
// User 用户
type User struct {
	Id         int    `json:"id" form:"id"`
	Username   string `json:"username" form:"username"`
	Password   string `json:"password" form:"password"`
	Status     int    `json:"status" form:"status"`
	CreateTime int64  `json:"createTime" form:"createTime"`
}

// 操作数据库
```

**Article**

```go
// Article 文章
type Article struct {
	Id             int
	Title          string
	Tags           string
	Short          string
	Content        string
	Author         string
	CreateTime     int64
	LastUpdateTime int64
}

// 操作数据库
```

# service 层

service 层算是 controller 和 model 的一个中间层，如果 model 是用来封装对象，处理对象最基本数据库操作的话，service 则是基于业务对 model 数据库再做了一层封装，并且被 controller 层调用。此项目中对 controller 需要封装的一些操作数据的业务属性操作进行了进一步封装并写在 service 层

项目中 service 没有和 controller 一一对应存在，应为有的 controller 中实际没有做 service 的封装，所以没有对应的 service。项目中 service 包中的 go 文件统一用 service 作为后缀

# controller 层

controller 层分为了好几个文件

- article 文章
- home 主页面
- login 登录页
- personal 关于我
- register 注册页
- session 处理登录态

**GET**

controller 中主要写的是访问路由绑定的各个 Controller 函数，其中有 GET 和 POST 类型的 controller 函数。其中对于 GET 类型函数多以 HTML 函数来渲染，如下

```go
// HomeGet 首页
func HomeGet(c *gin.Context) {
  // ...
  
  // 返回 html
	resp := gin.H{
		"isLogin":  isLogin,
		"username": GetSession(c, "loginUser"),
		"artNum":   artNum,
	}
	c.HTML(http.StatusOK, "home.html", resp)
}
```

**POST**

对于 POST 类型函数多以 json 形式来进行响应

```go
// LoginPost 登录请求
func LoginPost(c *gin.Context) {
  // ...
  
  // json 响应
  resp = gin.H{
    "code":    0,
    "message": "登录成功，欢迎进入！",
  }
  c.JSON(http.StatusOK, resp)
}
```

# view/static 前端

**view 中模版界定符**

因为前端 html 页面较多，所以在 view 包下进行了以下文件夹分类（依据页面模块分类），view 中存放的都是 html 页面，html 中使用了 go 的模版界定符，类似 JSP，同时也适用了 bootstrap v3框架等。虽然项目中使用了比较过时的 bootstrap 和 jquery 前端框架组建，但是对于 vue/react 大肆流行的今天，这些接近废弃的前端知识还是可以了解学习的（学习成本不会太高）

使用了模版界定符号，如下：

```html
<!-- 导航条 -->
{{template "nav.html" .}}

<!-- 文章列表 -->
{{.listContent}}
```

**view 中 bootstrap v3**

也使用了 bootstrap v3 框架，项目中使用的是在线版本的，当然觉得感兴趣也可以使用高版本版本 bootstrap 框架，并且可以下载下来放在 static 中调用使用。下方为在线引用 bootstrap v3

```html
<!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css"
          integrity="sha384-HSMxcRTRxnN+Bdg0JdbxYKrThecOKuH5zCYotlSAcp1+c8xmyTe9GYg1l9a69psu" crossorigin="anonymous">
<!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap-theme.min.css"
      integrity="sha384-6pzBo3FDv/PJ8r2KRkGHifhEocL+1X2rVCTTkUfGk7/0pbek5mMa1upzvWbrUbOZ" crossorigin="anonymous">
<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"
        integrity="sha384-aJ21OjlMXNL5UyIl/XNwTMqvzeRMZH2w8c5cRVpzpU8Y5bApTppSuUkhZXN0VxHd"
        crossorigin="anonymous"></script>
```

**view 中 jquery 及插件**

项目也是用了 jquery 及其插件，比如使用了 jquery 的校验插件

```go
<!-- jq 在线压缩版 -->
<script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/jquery-validation@1.19.3/dist/jquery.validate.min.js"></script>
```

**static 中 css**

项目中的 css 作为统一 css，主要写的是校验的 css 样式，刚兴趣的同学可以把更多 html 通用的 css 样式放在其中

**static 中 img**

这里保存有项目中要使用的图片资源

**static 中 js**

其中主要写的是各个 html 页面的 js 操作，其实主要还是表单提交操作

# database 数据库

数据库着一层在 main 中被调用来初始化数据库，其中主要存在三类函数

**数据库信息**

数据库服务在本地，端口号用默认

- 数据库版本：mysql 5.7
- 用户名：root
- 密码：123456
- 地址是：localhost
- 端口号：3306（默认）
- 数据库名：gindemo
- 创建后存在数据表：user 和 article 表

**Init 初始化数据库主函数**

这里主要操作是数据库的连接和调用创建表的函数，这些创建表的函数，如果发现表已经存在，则不会创建表。连接数据库语句如下，下面有一些数据是封装在 config 中的，方便后续配置

```go
// InitMysql 初始化 mysql 数据库
func InitMysql() {
  if DB == nil {
		DB, _ = sql.Open("mysql", config.DB_USER+":"+config.DB_PWD+"@tcp("+config.DB_ADD+":"+config.DB_PORT+")/"+config.DB_NAME)
  }
}
```

**建表函数**

建表函数被 Init 初始化数据库总函数所调用，主要是创建 user 表的函数和创建 article 表的函数

**对最基本数据库操作的封装**

因为 golang 中 sql 包 select 语句可以直接查询，但是 insert/update/create/delete 语句则需要 Exec 函数去执行，所以对基本的 Exec 执行函数进行了封装，也对基本的查询语句做了最简单的封装，方便各个 model 来调用

# util 工具类

因为注册需要填写用户和密码，但是直接把密码明文存入数据库有安全风险，所以打算把拿到的用户密码进行 md5 加密之后（go sdk 提供 md5 包）再存入数据库中，所以 util 在项目中存在一个加密函数，当然有朋友感兴趣也可以在 util 工具包中封装更多有意思的函数

```go
// MD5 md5 算法加密
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
```

# config 配置文件

因为项目中存在许多常量，这些常量可以进行封装，来方便项目中许多文件代码的统一调整。目前项目中的常量如下：

```go
/*----------- 数据库 ------------*/

// DB_ADD 数据库地址
const DB_ADD = "localhost"

// DB_PORT 数据库端口
const DB_PORT = "3306"

// DB_NAME 数据库名字
const DB_NAME = "gindemo"

// DB_USER 数据库用户
const DB_USER = "root"

// DB_PWD 数据库密码
const DB_PWD = "123456"

/*----------- 文章分类 ------------*/

// PROGRAM 文章分类-程序博客
const PROGRAM = "程序博客"

// FINANCE 文章分类-金融理财
const FINANCE = "金融理财"

// SCIENCE 文章分类-自然科学
const SCIENCE = "自然科学"

// ART 文章分类-音乐艺术
const ART= "音乐艺术"

// SPORT 文章分类-体育竞技
const SPORT = "体育竞技"

/*----------- 文章页面 article_list ------------*/

// SIZE_PERPAGE 文章列表页面显示每页最大显示文章数
const SIZE_PERPAGE = 5

// CHAR_PERARTICLE 文章列表页面每篇文章最大显示字符数量
const CHAR_PERARTICLE = 500
```

# 其他文件

**`.gitignore`**

其中用来填写取消跟踪的文件，依据自己的情况来填写取消跟踪的文件

**`go.mod`**

其中还存在 go.sum。因为采用 go module 来管理项目，所以存在`go.mod`文件，但是如果要上传 github，这个`go.mod`也是需要传的，因为相当于 java 的 pom 可以告知使用者该用哪一个具体版本。使用 golang 的话建议直接创建 go module 项目，这样项目下就会直接存在`go.mod`文件，之后如果要引用什么包的时候，引用完后再`go tidy`即可。如果对于没有创建 go module 项目来说也可以通过 go mod 命令来创建`go.mod`文件

**`README.md`**

这个不说了，项目的说明文档

# 改进建议

- **前端美观优化**

  前端页面美观程度可以美化，比如展示效果，css 显示，js 显示等

- **前端框架优化**

  前端使用了比较老的 bootstrap v3 版本，而且栅栏系统使用的不够规范，可以考虑优化，或者直接采用 vue 或者 react 前端框架

- **后端逻辑优化**

  目前的项目系统更像是一个博客管理系统，提供的功能除了所有用户注册和登录之外，还可以写文章展示文章等，但是只要登录之后所有文章都可以编辑修改，却不论登录用户是谁，不论登录用户是不是博客文章等撰写者都是可以对所有文章修改的，所以这一点可以优化修改。此项目更像是一个博客的后端管理系统，注册的都是博客后台管理者，可以对所有博客进行编辑操作

- **数据库操作层**

  可以考虑使用 gorm 来操作数据库

- **代码优化**

  代码中依然存在许多冗杂部分，可以考虑进一步优化修改和封装

- **功能添加**

  目前博客项目功能还十分有限，可以考虑添加更多的功能，添加更多功能就以为可能要使用更多的工具框架组建


