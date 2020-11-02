
<!-- Desgin -->
# Proxy server

## Hierachy


```shell
init
    |_ bootstrap
        |_ apps
            |_ loader
            |_ api-register
            |_ app-controller
        |_ monitor
        |_ network

control
    |_ apps
        |_ run
        |_ stop
        |_ static
    |_ pipes
        |_ msg
        |_ api
        |_ registery
        |_ middleware

run
    |_ control
        |_ apps
            |_ run
    |_ network
    |_ pipes
```

- Nework - (pipe) - App
- APP - (pipe) - APP
- APP API 1 - (pipe) - APP API 2


### bootstrap


#### 

> 框架设计和实现思路：
> - 需要解决的问题
> - 性能期望
> - 所面向的开发群体


1. 需要解决的问题
- 能快速实现业务功能，动态实现app加载和管理
- 封装通信问题，提供简单的通信方案，底层可实现多种通信方式包括不限于：RPC，http/https
- 封装监控，提供插件方案
- 框架是否可持续发展

> App： 一种服务的provider，同时也可能是其他服务的consumer

1.1 设计解决的问题

- 设计如何区分框架模块和App
  - 是否需要区分框架模块和App？
- 设计如何启动顺序，解决App项目依赖问题
  - 框架模块启动顺序
  - 如何有序启动依赖App
  - 如何设计App启动描述结构体
- 设计如何App静态配置
  - 如何传参
- 设计如何控制App
  - 生命周期： load config、 init、 pending、 running、closed
  - 操作：               load， init，     run，    close

（相关问题）

- 如何解决App之间数据传递
  - 数据类型
  - 反射，动态识别
- 如何设计内部接口参数传递和返回结果？
  - 直接调用，通过中间人转发调用，中间件插件放置在中间人
  - 通过chan做中间转发，共有上下行2个chan传输输入和输出，中间件插件放置在 chans 结构体中。
- 如何设计App提供多个接口同时能简便的暴露
  - 服务结构体存放Api map
  - 服务结构体提供Api map描述方法
  - 如何向外暴露Restful接口
  - 如何实现gRPC接口
  > http、gRPC绑定在一起


1.2 基本思路

- 整体生命周期分为： 加载配置、