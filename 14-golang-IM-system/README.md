
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

