### im-services
  * 简单的配置即可生成一个高性能、高可靠的消息推送服务器

#### 安装依赖
  * mysql
  * redis 
  * nsq(消息中间件)
  
#### 安装 nsq 
  ```shell
  docker-compose up -d //启动容器
  docker-compose ps //查看容器是否启动
  ```

#### 启动项目
```shell
 go run main.go 或者 air
```