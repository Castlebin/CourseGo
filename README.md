# 手把手，带你从零封装 Gin 框架

https://juejin.cn/post/7016742808560074783

https://github.com/jassue/jassue-gin

**注意，这里的项目名称为 jassue-gin，我们自己建的项目名字为 CourseGo。遇到 jassue-gin 的地方，需要替换为 CourseGo。**

## 1. 开篇 & 项目初始化
### 1.1. 项目初始化
```
mkdir CourseGo
cd CourseGo
go mod init CourseGo
```

然后加入 .gitignore 文件、README.md 文件。
接着执行 `git init` 命令，初始化项目为一个 git 仓库。

### 1.2. 加入 Gin 框架
```
go get -u github.com/gin-gonic/gin
```

### 1.3. 在项目根目录下编写 main.go 文件。测试下 gin 框架是否正常运行。

执行 `go run main.go` 命令，看看是否正常运行。
访问 http://127.0.0.1:8080/ping ，看看是否返回 pong。

### 1.4. 项目结构设计
| 文件/目录名称         | 	说明                   |
|-----------------|-----------------------|
| app/common      | 	公共模块（请求、响应结构体等）      |
| app/controllers | 	业务调度器                |
| app/middleware  | 	中间件                  |
| app/models      | 	数据库结构体               |
| app/services    | 	业务层                  |
| bootstrap       | 	项目启动初始化              |
| config          | 	配置结构体                |
| global          | 	全局变量                 |
| routes          | 	路由定义                 |
| static          | 	静态资源（允许外部访问）         |
| storage         | 	系统日志、文件等静态资源）        |
| utils           | 	工具函数                 |
| config.yaml     | 	配置文件                 |
| main.go         | 	项目启动文件               |



## 2. 配置初始化 & 全局变量
### 2.1. viper
我们使用 viper 来读取配置文件。首先安装 viper。
```
go get -u github.com/spf13/viper 
```
在项目根目录下新建一个文件 `config.yaml` ，初期先将项目的基本配置放入，后续我们会添加更多配置信息

### 2.2. 编写配置结构体
编写配置结构体
在项目根目录下新建文件夹 config，用于存放所有配置对应的结构体。进入该目录：   

新建 config.go 文件，定义 Configuration 结构体，其 App 属性对应 config.yaml 中的 app 属性  

新建 app.go 文件，定义 App 结构体，其所有属性分别对应 config.yaml 中 app 下的所有配置 

### 2.3. 全局变量
新建 global/app.go 文件，定义 Application 结构体，用来存放一些项目启动时的变量，便于调用，目前先将 viper 结构体和 Configuration 结构体放入，后续会添加其他成员属性

### 2.4. 使用 viper 载入配置
新建 bootstrap/config.go 文件，编写代码，使用 viper 载入配置文件

### 2.5. 在 main.go 中初始化配置
在 main.go 中初始化配置，将配置信息存入全局变量中。并且修改端口号从配置文件中读取

### 2.6. 测试
执行 `go run main.go` ，启动应用，可以看到服务器监听的端口是已经是配置文件里的端口号了。

