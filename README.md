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


## 3. 日志初始化
### 3.1. zap 和 lumberjack
我们使用 zap 来记录日志，使用 lumberjack 来切割归档日志文件。首先安装 zap 和 lumberjack。
```
go get -u go.uber.org/zap
go get -u gopkg.in/natefinch/lumberjack.v2
```

### 3.2. 编写 日志 的配置结构体
新建 config/log.go 文件，定义 zap 和 lumberjack 初始化需要使用的配置项，大家可以根据自己的喜好去定制

接下来别忘了在 config/config.go 文件中引入 Log 结构体。（记住，这个文件是 config.yaml 数据结构定义。所以，所有的配置项都应该加入到其中）

接着就可以在 config.yaml 增加对应配置项了。

### 3.3. 定义 utils 工具函数
新建 utils/directory.go 文件，编写 PathExists 函数，用于判断路径是否存在

### 3.4. 编写日志初始化代码
新建 bootstrap/log.go 文件，编写代码，初始化日志 

可以看到 跟我们初始化配置文件的方式是一样的。  
所以，接着也需要在 global/app.go 中，添加 Log 成员属性。并且在 main.go 中添加初始化日志的代码。
同样的，为了使日志功能方便调用，也将日志实例存入全局变量中 ( global/app.go )。

### 3.5. 测试
启动 main.go ，生成 storage/logs/app.log 文件，表示日志初始化成功


## 4. 数据库初始化（GORM)
### 4.1. GORM
我们使用 GORM 来操作数据库。首先安装 GORM。我们在这里使用 MySQL 数据库，所以还需要安装对应的 MySQL 驱动。
```
go get -u gorm.io/gorm

# GORM 官方支持 sqlite、mysql、postgres、sqlserver
go get -u gorm.io/driver/mysql 
```

### 4.2. 编写数据库配置结构体
新建 config/database.go 文件，自定义配置项，用于初始化数据库  

之后别忘了在 config/config.go 文件中引入 Database 结构体。  

接着就可以在 config.yaml 增加对应配置项了。  

### 4.3. 自定义 Logger（使用文件记录日志）
gorm 有一个默认的 logger ，由于日志内容是输出到控制台的，我们需要自定义一个写入器，将默认logger.Writer 接口的实现切换为自定义的写入器，上一篇引入了 lumberjack ，将继续使用它

新建 bootstrap/db.go 文件，编写 getGormLogWriter 函数

接下来，编写 getGormLogger 函数， 切换默认 Logger 使用的 Writer

至此，自定义 Logger 就已经实现了，这里只简单替换了 logger.Writer 的实现，大家可以根据各自的需求做其它定制化配置  


### 4.4.初始化数据库
在 bootstrap/db.go 文件中，编写 InitializeDB 初始化数据库函数，以便于在 main.go 中调用

### 4.5. 编写模型文件进行数据库迁移 
新建 app/models/common.go 文件，定义公用的数据库表模型字段  

新建 app/models/user.go 文件，定义 User 模型  

在 bootstrap/db.go 文件中，编写数据库表初始化代码。并且在 initMySqlGorm 函数中调用它。


### 4.6. 定义全局变量 DB
在 global/app.go 中，添加 DB 成员属性。并且在 main.go 中添加初始化数据库的代码。

### 4.7. 测试
由于我们使用了本地的 MySQL 数据库，所以需要先启动 MySQL 服务。并且，建好 go-dev 库。

启动 main.go ，可以看到数据库表 users 已经自动创建成功了。


## 5. 静态资源处理 & 优雅重启服务器
这一篇将对路由进行分组调整，把定义路由的文件集中到同一个目录下，并处理前端项目打包后的静态文件。在 Go 1.8 及以上版本中，内置的 http.Server 提供了 Shutdown() 方法，支持平滑重启服务器，本次将使用它调整项目启动代码，若 Go 版本低于 1.8 可以使用 fvbock/endless 来替代

### 5.1. 路由分组调整
新建 routes/api.go 文件，用来存放 api 分组路由 

新建 bootstrap/router.go 文件，用于项目启动时初始化路由。以及启动服务器，将 main.go 中的代码迁移过来  

在 main.go 文件中调用 RunServer() 方法，启动服务器  


### 5.2. 处理静态资源
新建 static 目录，用于存放前端项目打包后的静态资源。github 代码库中复制本工程的 static 目录下的文件到这里 

在 bootstrap/router.go 文件，setupRouter() 方法中编写相关静态资源的路由代码  

启动 main.go ，访问 http://localhost:8888/ ，前端资源处理成功    


### 5.3. 优雅重启/停止服务器
在 bootstrap/router.go 文件中，调整 RunServer() 方法。加入优雅停机的代码  

在 routes/api.go 中，添加一条测试路由 /test  

启动 main.go，访问 http://localhost:8888/api/test ，然后立即使用 CTRL + C 停止服务器。可以看到服务器接收到中止命令后，依旧等待 /api/test 接口完成响应后才停止服务器

## 6. 初始化 Validator & 封装 Response & 实现第一个接口
Gin 自带验证器返回的错误信息格式不太友好，本篇将进行调整，实现自定义错误信息，并规范接口返回的数据格式，分别为每种类型的错误定义错误码，前端可以根据对应的错误码实现后续不同的逻辑操作，篇末会使用自定义的 Validator 和 Response 实现第一个接口

### 6.1. 自定义验证器错误信息
新建 app/common/request/validator.go 文件，编写自定义验证器相关接口、类型定义和方法  

新建 app/common/request/user.go 文件，用来存放所有用户相关的请求结构体，并实现 Validator 接口  

在 routes/api.go 中编写测试代码  

启动服务器，使用 Postman 测试，自定义错误信息成功返回
```shell
curl --location 'http://localhost:8888/api/user/register' --header 'Content-Type: application/json' --data '{
    "name": "张三"
}'
```

### 6.2. 自定义验证器
有一些验证规则在 Gin 框架中是没有的，这个时候我们就需要自定义验证器

新建 utils/validator.go 文件，定义验证规则，后续有其他的验证规则将统一存放在这里 

新建 bootstrap/validator.go 文件，定制 Gin 框架 Validator 的属性  

在 main.go 中调用 bootstrap.InitializeValidator() 方法，初始化验证器

注：由于在 InitializeValidator() 方法中，使用 RegisterTagNameFunc() 注册了自定义 json tag， 所以在 GetMessages() 中自定义错误信息 key 值时，需使用 json tag 名称

重启服务器，使用 PostMan 测试，自定义验证器成功
```
curl --location 'http://localhost:8888/api/user/register' --header 'Content-Type: application/json' --data '{
"name": "张三",
"mobile": "12345678"
}'
```

返回结果：
```json
{"error":"手机号码格式不正确"}
```

### 6.3. 自定义错误码
新建 global/error.go 文件，将项目中可能存在的错误都统一存放到这里，为每一种类型错误都定义一个错误码，便于在开发过程快速定位错误，前端也可以根据不同错误码实现不同逻辑的页面交互

### 6.4. 封装 Response
新建 app/common/response/response.go 文件，定义 Response 结构体，用于规范接口返回的数据格式

### 6.5. 实现用户注册接口 
新建 utils/bcrypt.go 文件，编写密码加密及验证密码的方法   

新建 app/services/user.go 文件，编写用户注册逻辑  

新建 app/controllers/app/user.go 文件，校验入参，调用 UserService 注册逻辑  

在 routes/api.go 中，添加路由 `router.POST("/auth/register", app.Register)`

使用 Postman 调用接口 http://localhost:8888/api/auth/register ，如下图所示，接口返回成功，用户注册成功
![用户注册成功](doc-resourse/user-register-success.png)

```shell
curl --location 'http://localhost:8888/api/auth/register' --header 'Content-Type: application/json' --data '{
		    "name": "张三",
		    "mobile": "18912345678",
		    "password": "123456"
		}'
```

查看数据库 users 表，数据已成功写入


