
# tag操作


- 1. 在相应的本地仓库下打开git终端输入
```
git tag v1.0.0
```

- 1.1 tag与提交绑定


```
darren@darrendeMacBook-Pro 14-golang-IM-system % git commit -m "在线用户"
[test 68fc2d0] 在线用户
 4 files changed, 129 insertions(+), 2 deletions(-)
 create mode 100644 14-golang-IM-system/user.go



git tag V0.2 68fc2d0
```

- 2. 查看是否成功
```
git tag
```

- 3. 提交到远程
```
git push origin v1.0.0
```


- 1.删除本地tag
```
git tag -d tagName
```
- 2.删除远端tag
```
git push origin :refs/tags/tagName
```

- 3.创建本地tag
```
git tag tagName
```

- 4.提交到远端
```
git push origin --tags //提交全部tag
或者
git push origin tagName //提交单个tag
```
 
- 如果本地没有代码仓库：

```
git clone <git项目的地址>
git tag
git checkout <tag-name>

```



# 即实系统

## 服务端

```
 go build -o server server.go main.go user.go


  ./server 
```


## 客户端


```
nc 127.0.0.1 8888
```

- 编译客户端

```
go build -o client client.go 
```



## Go生态

- web框架 
  - beego  
  - gin
  - echo 
  - Iris 重量级的
- 微服务框架
  - Go kit 
  - Istio 
- 容器编排
  - k8s
  - swarm
- 服务发现
  - consul
- 存储引擎
  - k/v存储
  - 分布式存储 tidb
- 静态建站
  - hugo 
- 中间件
  - 消息队列 nsq
  - TCP长连接 zinx  
  - leaf(游戏服务器)
  - RPC   gRPC   
  - redis  codis
- 爬虫框架
  - go query 爬虫效率比python高
  