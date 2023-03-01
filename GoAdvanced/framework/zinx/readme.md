
![image.png](https://s2.loli.net/2023/02/07/CvNuKVOGreck3qW.png)


- V0.1 基本的server 
  * 方法
    + 启动服务器   ✅
    + 停止服务器   ✅
    + 运行服务器   ✅
      - 调用start()方法，调用之后阻塞处理
  * 属性
    + name名称   ✅
    + 监听的ip    ✅
    + 监听的端口  ✅
- V0.2 简单的链接封装和业务绑定
  * 链接的模块
    - 方法
      * 启动链接 Start()
      * 停止链接 Stop()
      * 获取当前链接的Conn对象
      * 得到链接ID
      * 得到客户端链接的地址和端口
      * 发送数据的方法Send()
    - 属性
      * socket TCP 套接字
      * 链接的ID
      * 当前连接的状态（是否已经关闭）
      * 与当前连接所绑定的处理业务方法
      * 等待被告知的channel 告知当前链接是退出的
- V0.3 基础route模块
  * Request 请求封装
    + 将链接和数据绑定在一起
      - 数据
        * 连接IConnection
        * 请求数据
      - 属性
        * 得到当前数据
        * 新建一个request请求
  * Router 模块
    + 抽象的IRouter
      - 处理业务之前的方法
      - 处理业务的主方法
      - 处理业务之后的方法
    + 具体的BaseRouter 
      - 处理业务之前的方法
      - 处理业务的主方法
      - 处理业务之后的方法
  * 集成Router模块
    + IServer添加路由添加功能
    + Server类添加Router成员
    + Connection绑定一个路由成员
    + 在Connection 调用已经注册的Router处理业务
      