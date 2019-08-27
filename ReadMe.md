# GoSpider 轻量级异步爬虫框架
`author：vophan`

编码不易，Star me！

## what's the GoSpider?
GoSpider的初衷其实是一个Go语言的入门级项目，大佬说，Golang三大入门项目`爬虫`，`博客`，`电商`。因为我是python转go，所以写一个爬虫真的是很小儿科了，所以，我想既然要做，不如去写一个框架，虽然之前用过`scrapy`这样的框架，但是，从来没有机会去自己想想框架是如何实现的，所以就有了GoSpider。

**但是，这并不代表他只是一个入门学习项目**，虽然是抱着学习的心态去做这个框架，但是，并没有想要放弃这个项目，我会维护这个项目，并不断更新，也希望有更多的朋友，小白，大佬一起加入进来。

## Structure

```
--GoSpider----conf				配置文件解析
		   |--demo				框架使用实例
		   |--spider			爬虫核心代码
		   |--test				测试代码
		   |--userAgent			userAgent文件
		   |--log				日志文件
		   |--utils				工具
```

框架的设计架构借鉴了`PSpider`的架构：

![](https://raw.githubusercontent.com/xianhu/PSpider/master/procedure.png)

## How to use it

目前，如果想要使用它，只需要：

> 1. 在根目录创建`main.go`
> 2. 在`main.go`中重写`parser`的`MiddleLinkRule`以及`LinkExtract`函数，`Saver`的`saveRule`函数
> 3. 根据需要修改`settings.json`配置文件

运行！

## For Learner

如果，你也像我一样，想要通过这个项目，熟悉Go语言，那么，实现它，**你将**：

1. 熟悉Golang的面向对象，接口，所谓的继承与多态，实际上的组合与接口
2. 了解Golang中的gorooutine，channel，sync，select等有关并发与协程的操作
3. 了解Golang中的http的相关操作
4. 熟悉Golang的语法

............

## For Developer

欢迎每一个想要加入的朋友，一起交流，毕竟他还是第一版，所以肯定大大小小的不足，所以，**我需要你们**：

1. 爬虫框架的log模块尚未实现
2. 爬虫框架的proxier模块尚需完善

.........

总之，我很菜，需要大家的帮助

## Contact Me

`qq`:`809866729`

`wx`:

![](https://raw.githubusercontent.com/skyf0cker/NumPy100/master/qrcode.png)