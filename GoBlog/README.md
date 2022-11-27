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

curl -X POST http://127.0.0.1:8000/api/v1/tags -F "name=GO11" -F "created_by=demo1"

curl -X POST 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2'

curl -X GET 'http://127.0.0.1:8000/api/v1/tags?page=1&page_size=2'

// 修改
curl -X PUT http://127.0.0.1:8000/api/v1/tags/2 -F state=0  -F "modified_by=darren"

curl -X DELETE http://127.0.0.1:8000/api/v1/tags/1