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

go install github.com/swaggo/swag/cmd/swag@v1.6.5


- validator参数校验

go get -u github.com/go-playground/validator/v10



# 接口

- 创tag

curl -X GET http://127.0.0.1:8000/api/v1/tags?state=6

