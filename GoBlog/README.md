# 博客项目

## 下载资源
- 配置文件
go get -u github.com/spf13/viper@v1.4.0

- 数据库
go get -u github.com/jinzhu/gorm@v1.9.12

- 日志
go get -u gopkg.in/natefinch/lumberjack.v2

- 接口文档

go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
go get -u github.com/swaggo/gin-swagger@v1.2.0
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template

- 下载swagger 
go install github.com/swaggo/swag/cmd/swag@v1.6.5


- validator参数校验

go get -u github.com/go-playground/validator/v10

- jwt
go get -u github.com/dgrijalva/jwt-go@v3.2.0


- email 
go get -u gopkg.in/gomail.v2

- 限流

go get -u github.com/juju/ratelimit@v1.0.1


- jaeger的使用

docker  run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
jaegertracing/all-in-one:1.16

5778 jaeger服务端口
16686 jaeger WebUI
9411 兼容zipkin的http端口



- 安装jaeger第三方库

go get -u github.com/opentracing/opentracing-go@v1.1.0
go get -u github.com/uber/jaeger-client-go@v2.22.1


- sql追踪
go get -u github.com/eddycjy/opentracing-gorm

# 接口

- 创tag

curl -X GET http://127.0.0.1:8000/api/v1/tags?state=6

curl -X POST http://127.0.0.1:8000/api/v1/tags -F "name=GO11" -F "created_by=demo1"

curl -X POST 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2'

curl -X GET 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2'



// 修改
curl -X PUT http://127.0.0.1:8000/api/v1/tags/2 -F state=0  -F "modified_by=darren"

curl -X DELETE http://127.0.0.1:8000/api/v1/tags/1


// 上传文件

curl -X POST http://127.0.0.1:8000/upload/file -F file=/Users/darren/go/src/github.com/darrenli6/GolangStudy/GoBlog/test.txt -F type=1 


# 启动问题

`go run main.go -port=8001 -mode=release -config=configs/`

- 打包进二进制文件中


`go get github.com/go-bindata/go-bindata/...`

`go install github.com/go-bindata/go-bindata/...`

可以配置文件打包进二进制文件中

`go-bindata -o configs/config.go -pkg=configs  configs/config.yaml`


## 配置热更新

```azure
go get -u golang.org/x/sys...
go get -u github.com/fsnotify/fsnotify    
```

## 优雅的重启
防止客户端访问中断


# gRPC

## 安装protoc编译器

`brew install protobuf`

`protoc --version`

## 安装protoc-gen-go插件

go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.3
go install  github.com/golang/protobuf/protoc-gen-go@v1.3.3

