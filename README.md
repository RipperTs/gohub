# gohub

## 接口文档
```
https://console-docs.apipost.cn/preview/95d2fdc25817753c/9c0003b1ee2e165e
```

## 程序结构
```
├── app                            // 程序具体逻辑代码
│   ├── cmd                         // 命令
│   │   ├── cache.go                
│   │   ├── cmd.go
│   │   ├── key.go
│   │   ├── make                    // make 命令及子命令
│   │   │   ├── make.go
│   │   │   ├── make_apicontroller.go
│   │   │   ├── make_cmd.go
│   │   │   ├── make_factory.go
│   │   │   ├── make_migration.go
│   │   │   ├── make_model.go
│   │   │   ├── make_policy.go
│   │   │   ├── make_request.go
│   │   │   ├── make_seeder.go
│   │   │   └── stubs               // make 命令的模板
│   │   │       ├── apicontroller.stub
│   │   │       ├── cmd.stub
│   │   │       ├── factory.stub
│   │   │       ├── migration.stub
│   │   │       ├── model
│   │   │       │   ├── model.stub
│   │   │       │   ├── model_hooks.stub
│   │   │       │   └── model_util.stub
│   │   │       ├── policy.stub
│   │   │       ├── request.stub
│   │   │       └── seeder.stub
│   │   ├── migrate.go
│   │   ├── play.go
│   │   ├── seed.go
│   │   └── serve.go
│   ├── http                        // http 请求处理逻辑
│   │   ├── controllers             // 控制器，存放 API 和视图控制器
│   │   │   ├── api                 // API 控制器，支持多版本的 API 控制器
│   │   │   │   └── v1              // v1 版本的 API 控制器
│   │   │   │       ├── users_controller.go
│   │   │   │       └── ...
│   │   └── middlewares             // 中间件
│   │       ├── auth_jwt.go
│   │       ├── guest_jwt.go
│   │       ├── limit.go
│   │       ├── logger.go
│   │       └── recovery.go
│   ├── models                      // 数据模型
│   │   ├── user                    // 单独的模型目录
│   │   │   ├── user_hooks.go       // 模型钩子文件
│   │   │   ├── user_model.go       // 模型主文件
│   │   │   └── user_util.go        // 模型辅助方法
│   │   └── ...
│   ├── policies                    // 授权策略目录
│   │   ├── category_policy.go
│   │   └── ...
│   └── requests                    // 请求验证目录（支持表单、标头、Raw JSON、URL Query）
│       ├── validators              // 自定的验证规则
│       │   ├── custom_rules.go
│       │   └── custom_validators.go
│       ├── user_request.go
│       └── ...
├── bootstrap                       // 程序模块初始化目录
│   ├── app.go  
│   ├── cache.go
│   ├── database.go
│   ├── logger.go
│   ├── redis.go
│   └── route.go
├── config                          // 配置信息目录
│   ├── app.go
│   ├── captcha.go
│   ├── config.go
│   ├── database.go
│   ├── jwt.go
│   ├── log.go
│   ├── mail.go
│   ├── pagination.go
│   ├── redis.go
│   ├── sms.go
│   └── verifycode.go
├── database                        // 数据库相关目录
│   ├── database.db                 // sqlite 数据文件（加入到 .gitignore 中）
│   ├── factories                   // 模型工厂目录
│   │   ├── user_factory.go
│   │   └── ...
│   ├── migrations                  // 数据库迁移目录
│   │   ├── 2021_12_21_102259_create_users_table.go
│   │   ├── 2021_12_21_102340_create_categories_table.go
│   │   └── ...
│   └── seeders                     // 数据库填充目录
│       ├── users_seeder.go
│       ├── ...
├── pkg                             // 内置辅助包
│   ├── app
│   ├── auth
│   ├── cache
│   ├── captcha
│   ├── config
│   └── ...
├── public                          // 静态文件存放目录
│   ├── css
│   ├── js
│   └── uploads                     // 用户上传文件目录
│       └── avatars                 // 用户上传头像目录
├── routes                          // 路由
│   ├── api.go
│   └── web.go
├── storage                         // 内部存储目录
│   ├── app
│   └── logs                        // 日志存储目录
│       ├── 2021-12-28.log
│       ├── 2021-12-29.log
│       ├── 2021-12-30.log
│       └── logs.log
└── tmp                             // air 的工作目录
├── .env                            // 环境变量文件
├── .env.example                    // 环境变量示例文件
├── .gitignore                      // git 配置文件
├── .air.toml                       // air 配置文件
├── .editorconfig                   // editorconfig 配置文件
├── go.mod                          // Go Module 依赖配置文件
├── go.sum                          // Go Module 模块版本锁定文件
├── main.go                         // Gohub 程序主入口
├── Makefile                        // 自动化命令文件
├── readme.md                       // 项目 readme
```

## 初始化项目
使用以下命令来初始化项目：  
```
go mod tidy
```
运行项目:  
```
go run main.go
```
推荐使用air命令来运行项目：
```
air
```
配置全局配置文件:  
```
将.env.example文件改名问.env 并在其内容中填写项目需要的配置信息。
```

## Air 自动重载
### 安装 air
使用以下命令来安装 air ：  
```
GO111MODULE=on  go install github.com/cosmtrek/air@latest  
```

安装成功后使用以下命令检查下：  
```
air -v

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ , built with Go
```

### 使用 air
在项目目录下运行以下命令：
```
air
```

### air 配置信息
我们可以使用 `.air.toml` 文件来配置 `air` 的行为。  


## 邮件服务器Mailhog安装
### 说明
发送邮件需要有邮件服务器，在开发『发送邮件验证码功能』前，我们先来创建邮件服务器  

#### Mailhog
命令行安装 Mailhog ：  
```
GO111MODULE=on  go install github.com/mailhog/MailHog@latest
```
安装完成后运行 Mailhog ：  
```
MailHog
```
接下来将会输出:  
```
2022/06/12 12:03:28 Using in-memory storage
2022/06/12 12:03:28 [SMTP] Binding to address: 0.0.0.0:1025
[HTTP] Binding to address: 0.0.0.0:8025
2022/06/12 12:03:28 Serving under http://0.0.0.0:8025/
Creating API v1 with WebPath: 
Creating API v2 with WebPath: 
```
以上输出有两个比较重要的信息：  

- 0.0.0.0:1025 是 SMTP 端口
- 0.0.0.0:8025/ 是网页界面  

#### Web 界面
访问 http://127.0.0.1:8025/ 