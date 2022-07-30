### im-services

[![OSCS Status](https://www.oscs1024.com/platform/badge/IM-Tools/Im-Services.svg?size=small)](https://www.oscs1024.com/project/IM-Tools/Im-Services?ref=badge_small)


  * 简单的配置即可生成一个高性能、高可靠的消息推送服务器

  * [docker-compose 安装](docs/1.安装使用.md)

#### 实现
  * 负载:实现了简易版集群(节点消息依靠Grpc传输消息)
  * 消息:实现了Ws消息投递、Api消息投递、Grpc消息投递。以及自定义消息




















  * 项目结构
```shell
   ─ app //应用核心
│   ├── api //接口
│   │   ├── controllers // Api控制器
│   │   ├── requests  // 接口请求校验
│   │   └── services //封装服务层
│   ├── dao //数据层
│   ├── enum    //枚举
│   ├── helpers //辅助函数
│   ├── middleware //中间件
│   ├── models // 数据库模型
│   ├── router // 路由
│   └── service // 网络层 核心逻辑层
│       ├── bootstrap // 启动服务
│       ├── cache // 缓存逻辑
│       ├── client // ws客户端
│       ├── dao  // 数据层
│       ├── dispatch // 服务节点
│       ├── handler // 控制器
│       ├── message // ws消息处理
│       ├── queue // 消息中间件
│       └── tests // 测试文件
├── cmd // cli 命令工具
│   ├── cmd
├── config // 配置文件加载
├── config.yaml // 配置文件
├── config.yaml.test
├── docker // docker配置文件 应用环境
├── docs // 项目文档
├── go.mod
├── go.sum
├── main.go // 入口文件
├── pkg // 第三方包封装
├── server // grpc服务端和客户端
│   ├── client //客户端方法
│   ├── grpc //根据protos生成的文件
│   ├── protos //定义的protos
│   ├── run.go //启动grpc入口文件
│   └── server // 服务层方法
├── storage // 日志以及静态文件
│   └── logs 

 ```